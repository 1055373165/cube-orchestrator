# cube-orchestrator

## Start And API Test

> start docker desktop Engine on your computer

### Service start

```
CUBE_WORKER_HOST=127.0.0.1 \
CUBE_WORKER_PORT=5555 \
CUBE_MANAGER_HOST=127.0.0.1 \
CUBE_MANAGER_PORT=5556 \
go run main.go
```
### POST new task

curl -v --request POST \         
--header 'Content-Type: application/json' \
--data @task1.json \
localhost:5556/tasks

```
2024/09/08 16:02:59 Pulled {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0} off pending queue
2024/09/08 16:02:59 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/08 16:02:59 task.Task{ID:uuid.UUID{0xbb, 0x1d, 0x59, 0xef, 0x9f, 0xc1, 0x4e, 0x4b, 0xa4, 0x4d, 0xdb, 0x57, 0x1e, 0xee, 0xd2, 0x3}, ContainerID:"", Name:"test-chapter-9.1", State:1, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet{"7777/tcp":struct {}{}}, HostPorts:nat.PortMap(nil), PortBindings:map[string]string{"7777/tcp":"7777"}, RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"/health", RestartCount:0}
2024/09/08 16:02:59 Sleeping for 10 seconds

2024/09/08 16:03:09 Performing task health check
2024/09/08 16:03:09 Task health checks completed
2024/09/08 16:03:09 Sleeping for 60 seconds

{"status":"Pulling from sun4965485/echo-smy","id":"v1"}
{"status":"Digest: sha256:b3a6951a31ab9ba821c95815ccc16de992fd00019fab37ed607514e61cf6f6fe"}
{"status":"Status: Image is up to date for sun4965485/echo-smy:v1"}
2024/09/08 16:03:12 task bb1d59ef-9fc1-4e4b-a44d-db571eeed203 Running on container d72016f22d531ce713fb1673a1cc6e7d52b77bc743774861afb0c3a094283333

2024/09/08 16:06:24 Checking worker 127.0.0.1:5555 for task updates
2024/09/08 16:06:24 Collecting stats
2024/09/08 16:06:24 Attempting to update task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
```
### Retrieve task list

curl -v localhost:5556/tasks | jq

```
[
  {
    "ID": "bb1d59ef-9fc1-4e4b-a44d-db571eeed203",
    "ContainerID": "d72016f22d531ce713fb1673a1cc6e7d52b77bc743774861afb0c3a094283333",
    "Name": "test-chapter-9.1",
    "State": 2,
    "Image": "sun4965485/echo-smy:v1",
    "CPU": 0,
    "Memory": 0,
    "Disk": 0,
    "ExposedPorts": {
      "7777/tcp": {}
    },
    "HostPorts": null,
    "PortBindings": {
      "7777/tcp": "7777"
    },
    "RestartPolicy": "",
    "StartTime": "0001-01-01T00:00:00Z",
    "FinishTime": "0001-01-01T00:00:00Z",
    "HealthCheck": "/health",
    "RestartCount": 0
  }
]
```

### Post task event to complete task
curl -v --request POST \
--header 'Content-Type: application/json' \
--data @stop_task.json \

```
2024/09/08 16:07:09 Pulled {bb1d59ef-9fc1-4e4b-a44d-db571eeed203 d72016f22d531ce713fb1673a1cc6e7d52b77bc743774861afb0c3a094283333 test-chapter-9.1 3 sun4965485/echo-smy:v1 0 0 0 map[] map[] map[]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC  0} off pending queue
2024/09/08 16:07:09 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/08 16:07:09 task.Task{ID:uuid.UUID{0xbb, 0x1d, 0x59, 0xef, 0x9f, 0xc1, 0x4e, 0x4b, 0xa4, 0x4d, 0xdb, 0x57, 0x1e, 0xee, 0xd2, 0x3}, ContainerID:"d72016f22d531ce713fb1673a1cc6e7d52b77bc743774861afb0c3a094283333", Name:"test-chapter-9.1", State:3, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet(nil), HostPorts:nat.PortMap(nil), PortBindings:map[string]string(nil), RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"", RestartCount:0}
2024/09/08 16:07:09 Sleeping for 10 seconds
2024/09/08 16:07:12 Atemmpting to stop container d72016f22d531ce713fb1673a1cc6e7d52b77bc743774861afb0c3a094283333
2024/09/08 16:07:12 Stopped and removed container d72016f22d531ce713fb1673a1cc6e7d52b77bc743774861afb0c3a094283333 for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/08 16:07:12 Sleeping 10 time seconds
```

### check container status

$ docker ps -a              
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES