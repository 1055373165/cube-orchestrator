# cube-orchestrator

## Start And API Test

> start docker desktop Engine on your computer

CUBE_WORKER_HOST=127.0.0.1 \
CUBE_WORKER_PORT=5555 \
CUBE_MANAGER_HOST=127.0.0.1 \
CUBE_MANAGER_PORT=5556 \
go run main.go

```
curl -v localhost:5556/tasks
```
[
  {
    "ID": "21b23589-5d2d-4731-b5c9-a97e9832d021",
    "ContainerID": "",
    "Name": "test-container-0",
    "State": 1,
    "Image": "containous/whoami",
    "CPU": 0,
    "Memory": 0,
    "Disk": 0,
    "ExposedPorts": null,
    "PortBindings": null,
    "RestartPolicy": "",
    "StartTime": "0001-01-01T00:00:00Z",
    "FinishTime": "0001-01-01T00:00:00Z"
  }
]

```
curl -v --request POST \         
--header 'Content-Type: application/json' \
--data @task.json \
localhost:5556/tasks
```

{"status":"Pulling from containous/whoami","id":"latest"}
{"status":"Digest: sha256:7d6a3c8f91470a23ef380320609ee6e69ac68d20bc804f3a1c6065fb56cfa34e"}
{"status":"Status: Image is up to date for containous/whoami:latest"}
2024/09/08 02:17:09 task 21b23589-5d2d-4731-b5c9-a97e9832d021 Running on container 1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169
2024/09/08 02:17:09 Sleeping 10 time seconds


```
curl -v localhost:5556/tasks
```
[
  {
    "ID": "21b23589-5d2d-4731-b5c9-a97e9832d021",
    "ContainerID": "1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169",
    "Name": "test-container-0",
    "State": 2,
    "Image": "containous/whoami",
    "CPU": 0,
    "Memory": 0,
    "Disk": 0,
    "ExposedPorts": null,
    "PortBindings": null,
    "RestartPolicy": "",
    "StartTime": "0001-01-01T00:00:00Z",
    "FinishTime": "0001-01-01T00:00:00Z"
  }
]


```
curl -v --request POST \
--header 'Content-Type: application/json' \
--data @stop_task.json \
localhost:5556/tasks
```

2024/09/08 02:21:59 Atemmpting to stop container 1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169
2024/09/08 02:21:59 Stopped and removed container 1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169 for task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/08 02:21:59 Sleeping 10 time seconds

[
  {
    "ID": "21b23589-5d2d-4731-b5c9-a97e9832d021",
    "ContainerID": "1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169",
    "Name": "test-container-0",
    "State": 3,
    "Image": "containous/whoami",
    "CPU": 0,
    "Memory": 0,
    "Disk": 0,
    "ExposedPorts": null,
    "PortBindings": null,
    "RestartPolicy": "",
    "StartTime": "0001-01-01T00:00:00Z",
    "FinishTime": "2024-09-07T18:21:59.378513698Z"
  }
]


$ docker ps -a
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES