package task

import (
	"context"
	"io"
	"log"
	"math"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	ContainerID string
	Name        string
	State       State
	Image       string
	CPU         float64
	Memory      int64
	Disk        int64
	// 这是容器镜像中声明的端口，表示容器内部暴露的端口。例如，"7777/tcp": {} 表示容器内部的 7777 端口使用 TCP 协议。
	ExposedPorts nat.PortSet
	// 这是 Docker 容器启动时的端口绑定配置，表示将容器内部的端口映射到主机上的端口。例如，"7777/tcp": "7777" 表示将容器内部的 7777 端口映射到主机上的 7777 端口。
	HostPorts     nat.PortMap
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
	HealthCheck   string
	RestartCount  int
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

// Config struct to hold Docker container config
type Config struct {
	// Name of the task, also used as the container name
	Name string // "test-container-1"
	// AttachStdin boolean which determines if stdin should be attached
	AttachStdin bool
	// AttachStdout boolean which determines if stdout should be attached
	AttachStdout bool
	// AttachStderr boolean which determines if stderr shoulb be attached
	AttachStderr bool
	// ExposedPorts list of posts exposed
	ExposedPorts nat.PortSet // Port is a string containing port number and protocol in the format "80/tcp""tcp:80".
	// Cmd to be run inside container (optional)
	Cmd []string
	// Image used to run the container
	Image string
	// CPU
	CPU float64
	// Memory in MiB
	Memory int64
	// Disk in GiB
	Disk int64
	// Env variables
	Env []string // Allows a user to specify environment variables that will get passed into the container.
	// RestartPolicy for the container ["", "always", "unless-stopped", "on-failure"]
	RestartPolicy string // Tells the Docker daemon what to do if a container dies unexpectedly.
}

func NewConfig(t *Task) *Config {
	return &Config{
		Name:          t.Name,
		ExposedPorts:  t.ExposedPorts,
		Image:         t.Image,
		CPU:           t.CPU,
		Memory:        t.Memory,
		Disk:          t.Disk,
		RestartPolicy: t.RestartPolicy,
	}
}

type Docker struct {
	Client *client.Client
	Config Config
}

func NewDocker(c *Config) *Docker {
	// NewClientWithOpts initializes a new API client with a default HTTPClient,
	// and default API host and version. It also initializes the custom HTTP headers to add to each request.
	dc, _ := client.NewClientWithOpts(client.FromEnv)
	return &Docker{
		Client: dc,
		Config: *c,
	}
}

type DockerResult struct {
	Error       error
	Action      string // start or stop
	ContainerId string
	Result      string
}

type DockerInspectResponse struct {
	Error     error
	Container *types.ContainerJSON
}

func (d *Docker) Run(portBinds map[string]string) DockerResult {
	ctx := context.Background()
	// Pull the image
	reader, err := d.Client.ImagePull(
		ctx, d.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	io.Copy(os.Stdout, reader)

	// Set restart policy
	rp := container.RestartPolicy{
		Name: container.RestartPolicyMode(d.Config.RestartPolicy),
	}

	// Set resource limits
	r := container.Resources{
		Memory:   d.Config.Memory,
		NanoCPUs: int64(d.Config.CPU * math.Pow(10, 9)), // CPU quota in units of 10<sup>-9</sup> CPUs.
	}

	// Set exposed ports
	exposedPorts := nat.PortSet{}
	for port := range d.Config.ExposedPorts {
		exposedPorts[nat.Port(port)] = struct{}{}
	}

	// Set container config
	cc := container.Config{
		Image:        d.Config.Image,
		Tty:          false,
		Env:          d.Config.Env,
		ExposedPorts: d.Config.ExposedPorts,
	}

	// Set port bindings
	portBindings := nat.PortMap{}
	// "7777/tcp": "7778"
	for containerPort, hostPort := range portBinds {
		portBindings[nat.Port(containerPort)] = []nat.PortBinding{
			{
				HostIP:   "127.0.0.1",
				HostPort: hostPort,
			},
		}
	}

	// Set host config
	hc := container.HostConfig{
		RestartPolicy:   rp,
		Resources:       r,
		PublishAllPorts: true,
		PortBindings:    portBindings,
	}

	// Create the container
	resp, err := d.Client.ContainerCreate(ctx, &cc, &hc, nil, nil, d.Config.Name)
	if err != nil {
		log.Printf("Error creating container using image %s: %v\n",
			d.Config.Image, err)
		return DockerResult{Error: err}
	}

	// Start the container
	err = d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		log.Printf("Error starting container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	// Get container logs
	out, err := d.Client.ContainerLogs(ctx, resp.ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return DockerResult{
		ContainerId: resp.ID,
		Action:      "start",
		Result:      "success",
	}
}

func (d *Docker) Stop(id string) DockerResult {
	log.Printf("Atemmpting to stop container %v", id)
	ctx := context.Background()
	err := d.Client.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		log.Printf("Error stopping containers %s: %v\n", id, err)
		return DockerResult{Error: err}
	}

	err = d.Client.ContainerRemove(ctx, id, container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		log.Printf("Error removing container %s: %v\n", id, err)
		return DockerResult{Error: err}
	}

	return DockerResult{
		Action: "stop",
		Result: "success",
		Error:  nil,
	}
}

func (d *Docker) Inspect(containerID string) DockerInspectResponse {
	dc, _ := client.NewClientWithOpts(client.FromEnv)
	resp, err := dc.ContainerInspect(context.Background(), containerID)
	if err != nil {
		log.Printf("Error inspecting container: %s\n", err)
		return DockerInspectResponse{Error: err}
	}
	return DockerInspectResponse{Container: &resp}
}
