# cube-orchestrator

cube-orchestrator is the core implementation of the container scheduler.

## Start And API Test

> start docker desktop Engine on your computer

### Service start

```
CUBE_WORKER_HOST=127.0.0.1 \
CUBE_WORKER_PORT=5000 \
CUBE_MANAGER_HOST=127.0.0.1 \
CUBE_MANAGER_PORT=5556 \
go run main.go
```
### POST new task

#### Command
```
curl -v --request POST --header 'Content-Type: application/json' --data @task1.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @task2.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @task3.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @task4.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @stop_task.json http://127.0.0.1:5556/tasks
```

#### Log 
```
2024/09/11 07:38:09 Add event {a7aa1d44-08f6-443e-9378-f5884311018e 2 0001-01-01 00:00:00 +0000 UTC {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7778]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/11 07:38:09 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/11 07:38:09 Add event {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {21b23589-5d2d-4731-b5c9-a97e9832d021  test-chapter-9.2 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7779]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/11 07:38:09 Added task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/11 07:38:09 Add event {a7aa1d44-08f6-443e-9378-f5884311719e 2 0001-01-01 00:00:00 +0000 UTC {95fbe134-7f19-496a-acfc-c7853e5b4cd2  test-chapter-9.3 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7800]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/11 07:38:09 Added task 95fbe134-7f19-496a-acfc-c7853e5b4cd2
2024/09/11 07:38:09 Add event {a7aa1d44-08f6-443e-9378-f5864313419e 2 0001-01-01 00:00:00 +0000 UTC {95fbe134-7f19-496a-acfc-c7753e5b4cd2  test-chapter-9.4 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7801]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/11 07:38:09 Added task 95fbe134-7f19-496a-acfc-c7753e5b4cd2
2024/09/11 07:38:16 Processing any tasks in the queue
2024/09/11 07:38:16 Pulled {a7aa1d44-08f6-443e-9378-f5884311018e 2 0001-01-01 00:00:00 +0000 UTC {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7778]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queues
```

## All running tasks

### Log 
```
[worker] Found task in queue: {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7778]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}
{"status":"Pulling from sun4965485/echo-smy","id":"v1"}
{"status":"Digest: sha256:b3a6951a31ab9ba821c95815ccc16de992fd00019fab37ed607514e61cf6f6fe"}
{"status":"Status: Image is up to date for sun4965485/echo-smy:v1"}
2024/09/11 07:38:29 task bb1d59ef-9fc1-4e4b-a44d-db571eeed203 Running on container 1504351ad1d291c5ee50cfc19d9537268ad730df6d7c3ef4daeeb7e2b75c788b
```

### List task
```
* Connection #0 to host localhost left intact
[
    {
    "ID": "bb1d59ef-9fc1-4e4b-a44d-db571eeed203",
    "ContainerID": "1504351ad1d291c5ee50cfc19d9537268ad730df6d7c3ef4daeeb7e2b75c788b",
    "Name": "test-chapter-9.1",
    "State": 2, // Running 
    "Image": "sun4965485/echo-smy:v1",
    "CPU": 0,
    "Memory": 0,
    "Disk": 0,
    "ExposedPorts": {
      "7777/tcp": {}
    },
    "HostPorts": null,
    "PortBindings": {
      "7777/tcp": "7778"
    },
    "RestartPolicy": "",
    "StartTime": "0001-01-01T00:00:00Z",
    "FinishTime": "0001-01-01T00:00:00Z",
    "HealthCheck": "/health",
    "RestartCount": 0
  },
  ....
]
```

## Simulate task over

### construct stop_task.json

```
{
    "ID": "6be4cb6b-61d1-40cb-bc7b-9cacefefa60c",
    "State": 3,
    "Task": {
        "State": 2, 
        "ID": "bb1d59ef-9fc1-4e4b-a44d-db571eeed203", 
	    "ContainerID": "1504351ad1d291c5ee50cfc19d9537268ad730df6d7c3ef4daeeb7e2b75c788b",
        "Name": "test-chapter-9.1", 
        "Image": "sun4965485/echo-smy:v1"
    }
}
```

### Command

> curl -v --request POST --header 'Content-Type: application/json' --data @stop_task.json http://127.0.0.1:5556/tasks 

### task status

```
{
    "ID": "bb1d59ef-9fc1-4e4b-a44d-db571eeed203",
    "ContainerID": "1504351ad1d291c5ee50cfc19d9537268ad730df6d7c3ef4daeeb7e2b75c788b",
    "Name": "test-chapter-9.1",
    "State": 3, // Completed
    "Image": "sun4965485/echo-smy:v1",
    "CPU": 0,
    "Memory": 0,
    "Disk": 0,
    "ExposedPorts": {
      "7777/tcp": {}
    },
    ...
}
```


## same container run same port and mapping different host port

```
CONTAINER ID   IMAGE                    COMMAND       CREATED         STATUS         PORTS                      NAMES
2cd1dd5d5cba   sun4965485/echo-smy:v1   "/app/echo"   4 minutes ago   Up 4 minutes   127.0.0.1:7801->7777/tcp   test-chapter-9.4
22c79c936ab3   sun4965485/echo-smy:v1   "/app/echo"   5 minutes ago   Up 5 minutes   127.0.0.1:7800->7777/tcp   test-chapter-9.3
ba718d096964   sun4965485/echo-smy:v1   "/app/echo"   5 minutes ago   Up 5 minutes   127.0.0.1:7779->7777/tcp   test-chapter-9.2
232d41465d9d   sun4965485/echo-smy:v1   "/app/echo"   5 minutes ago   Up 5 minutes   127.0.0.1:7778->7777/tcp   test-chapter-9.1
```

## update check for every worker on every task
```
2024/09/09 15:31:40 Checking worker 127.0.0.1:5000 for task updates
2024/09/09 15:31:40 Checking worker 127.0.0.1:5001 for task updates
2024/09/09 15:31:40 Checking worker 127.0.0.1:5002 for task updates
```

## EVPM Schduler Algorithm for load balancing
```
2024/09/09 15:32:30 [manager] selected worker 127.0.0.1:5000 for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/09 15:32:40 [manager] selected worker 127.0.0.1:5000 for task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/09 15:33:00 [manager] selected worker 127.0.0.1:5001 for task 95fbe134-7f19-496a-acfc-c7753e5b4cd2
```

## Task include all info to run container
```
2024/09/09 15:39:57 task bb1d59ef-9fc1-4e4b-a44d-db571eeed203 Running on container 5f38f532361917feb12b1c2071799c02e2cba47895673b5183d02cc27a8a7eb1
2024/09/09 15:39:57 task 21b23589-5d2d-4731-b5c9-a97e9832d021 Running on container 9b1e2f405e4b2813cf8d39442d212bb2551ffa7e3772c11735840e9f2841d942
...
```

## Health checks for every worker on every task
```
2024/09/11 01:15:09 Calling health check for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203: /health
2024/09/11 01:15:09 Calling health check for worker localhost task bb1d59ef-9fc1-4e4b-a44d-db571eeed203: http://localhost:7778/health
2024/09/11 01:15:09 Task bb1d59ef-9fc1-4e4b-a44d-db571eeed203 health check response: 200
2024/09/11 01:15:09 Calling health check for task 21b23589-5d2d-4731-b5c9-a97e9832d021: /health
2024/09/11 01:15:09 Calling health check for worker localhost task 21b23589-5d2d-4731-b5c9-a97e9832d021: http://localhost:7779/health
2024/09/11 01:15:09 Task 21b23589-5d2d-4731-b5c9-a97e9832d021 health check response: 200
2024/09/11 01:15:09 Calling health check for task 95fbe134-7f19-496a-acfc-c7853e5b4cd2: /health
2024/09/11 01:15:09 Calling health check for worker localhost task 95fbe134-7f19-496a-acfc-c7853e5b4cd2: http://localhost:7800/health
2024/09/11 01:15:09 Task 95fbe134-7f19-496a-acfc-c7853e5b4cd2 health check response: 200
2024/09/11 01:15:09 Calling health check for task 95fbe134-7f19-496a-acfc-c7753e5b4cd2: /health
2024/09/11 01:15:09 Calling health check for worker localhost task 95fbe134-7f19-496a-acfc-c7753e5b4cd2: http://localhost:7801/health
2024/09/11 01:15:09 Task 95fbe134-7f19-496a-acfc-c7753e5b4cd2 health check response: 200
2024/09/11 07:00:02 Calling health check for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203: /health
2024/09/11 07:00:02 Calling health check for worker localhost task bb1d59ef-9fc1-4e4b-a44d-db571eeed203: http://localhost:7778/health
```

## Collect container port binding info
```
2024/09/11 06:59:52 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7778"}}}
```

### Collect  all worker node stats about cpu, disk and memory which will be used to scedule

```
2024/09/11 07:26:56 collect stats from worker http://127.0.0.1:5000 success, CPU detail: linux.CPUStat{Id:cpu, User:338868, Nice:1029, System:107121, Idle:3540676, IOWait:7454, IRQ:0, SoftIRQ:4443, Steal:0, Guest:0, GuestNice:0}
2024/09/11 07:26:56 collect stats from worker http://127.0.0.1:5000 success, Disk detail: linux.Disk{All:66205626368, Used:17017856000, Free:49187770368, FreeInodes:3723900}
2024/09/11 07:26:56 collect stats from worker http://127.0.0.1:5000 success, Memory detail: linux.MemInfo{MemTotal:2014312, MemFree:93108, MemAvailable:536020, Buffers:23580, Cached:623660, SwapCached:54440, Active:812156, Inactive:683684, ActiveAnon:504656, InactiveAnon:529904, ActiveFile:307500, InactiveFile:153780, Unevictable:177560, Mlocked:80, SwapTotal:2097148, SwapFree:1146984, Dirty:536, Writeback:0, AnonPages:1008156, Mapped:204116, Shmem:185960, Slab:163100, SReclaimable:71744, SUnreclaim:91356, KernelStack:11008, PageTables:27076, NFS_Unstable:0, Bounce:0, WritebackTmp:0, CommitLimit:3104304, Committed_AS:7552064, VmallocTotal:133143592960, VmallocUsed:29268, VmallocChunk:0, HardwareCorrupted:0, AnonHugePages:0, HugePages_Total:0, HugePages_Free:0, HugePages_Rsvd:0, HugePages_Surp:0, Hugepagesize:2048, DirectMap4k:0, DirectMap2M:0, DirectMap1G:0}
```
