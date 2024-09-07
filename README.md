# cube-orchestrator

## Start And API Test

> start docker desktop Engine on your computer

export CUBE_HOST=localhost
export CUBE_PORT=5555
go run main.go

sudo apt install jq
curl -v http://127.0.0.1:5555/tasks | jq

curl -v --request POST \
--header 'Content-Type: application/json' \
--data @stop_task.json \
http://127.0.0.1:5555/tasks


stop_task.json: Task_ID = b07c14c8-9484-4174-b5ee-ebd7b7cfc359
stop_task.json: ContainerID = 1de683c4ab9c5d141b88e9520eba3fdb5f12f2a1e50ec7143d40ea458b28cfdd
stop_task.json: Task Now State=2
stop_task.json: Task Expected State=3 (Completed)

curl -v http://127.0.0.1:5555/tasks | jq

```
[
  {
    "ID": "b07c14c8-9484-4174-b5ee-ebd7b7cfc359",
    "ContainerID": "1de683c4ab9c5d141b88e9520eba3fdb5f12f2a1e50ec7143d40ea458b28cfdd",
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
    "FinishTime": "2024-09-07T15:53:59.003688063Z"
  },
  {
    "ID": "d1eb0067-ffd7-46af-ad78-eb24c3e79ecd",
    "ContainerID": "003c713c6a933f26bd5089ecf7d951255baa9c52af3a0f30c3242a992ca382ec",
    "Name": "test-container-1",
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
  },
  {
    "ID": "e86b8036-9ccb-497a-bbf8-078b9251c082",
    "ContainerID": "c97a71a8ebbad990cda952fe149f5f352c74a2bee7c602c2bfee20a3731bf313",
    "Name": "test-container-2",
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