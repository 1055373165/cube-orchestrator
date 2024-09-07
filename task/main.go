package task

import (
	"log"

	"github.com/docker/docker/client"
)

func main() {
	c := Config{
		Name:  "test-container-1",
		Image: "postgres:13",
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := Docker{
		Client: dc,
		Config: c,
	}

	log.Printf("docker client created: %v\n", dc)
	d.Run()
}
