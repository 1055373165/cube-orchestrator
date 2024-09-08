# cube-orchestrator

## Start And API Test

> start docker desktop Engine on your computer

1. Service start
```
CUBE_WORKER_HOST=127.0.0.1 \
CUBE_WORKER_PORT=5555 \
CUBE_MANAGER_HOST=127.0.0.1 \
CUBE_MANAGER_PORT=5556 \
go run main.go
```

2. Search task list
```
curl -v localhost:5556/tasks
```
Output:
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

3. POST new task to manager, manager distribute to worker
```
curl -v --request POST \         
--header 'Content-Type: application/json' \
--data @task.json \
localhost:5556/tasks
```
Output:
{"status":"Pulling from containous/whoami","id":"latest"}
{"status":"Digest: sha256:7d6a3c8f91470a23ef380320609ee6e69ac68d20bc804f3a1c6065fb56cfa34e"}
{"status":"Status: Image is up to date for containous/whoami:latest"}
2024/09/08 02:17:09 task 21b23589-5d2d-4731-b5c9-a97e9832d021 Running on container 1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169
2024/09/08 02:17:09 Sleeping 10 time seconds

4. Search task list again
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


5. Simulate task completed, stop and remove container which run task
```
curl -v --request POST \
--header 'Content-Type: application/json' \
--data @stop_task.json \
localhost:5556/tasks
```

Output:
2024/09/08 02:21:59 Atemmpting to stop container 1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169
2024/09/08 02:21:59 Stopped and removed container 1e06640f2f758780337551b0fb6af5819ce2ab79e192bacb98e55a8714c62169 for task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/08 02:21:59 Sleeping 10 time seconds

6. Search task list again

```
curl -v localhost:5556/tasks
```

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

8. check if the container has been stopped and removed
$ docker ps -a
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
