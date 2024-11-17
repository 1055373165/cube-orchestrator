# cube-orchestrator

cube-orchestrator is the core implementation of the container scheduler.

![image](https://github.com/user-attachments/assets/a4a3697a-4672-4823-9705-548513ad168e)

The system consists of three main components: the **Scheduler**, the **Manager**, and the **Worker**.

**Scheduler**

The **Scheduler** operates in three generic phases: **Feasibility Analysis**, **Scoring**, and **Picking**, executed sequentially as it assigns tasks to appropriate workers.

- **Feasibility Analysis**: This phase assesses whether it is possible to assign a task to a worker. In some cases, the task may not be assignable to any worker, while in others, it may only be assignable to a subset of workers. This phase can be compared to buying a car: with a budget of 100,000 RMB, the cars you can choose from depend on the dealership. Some dealerships may have only cars exceeding your budget, while others may have a subset of cars that fit your budget.

- **Scoring**: In this phase, the scheduler scores the workers identified in the feasibility analysis based on certain criteria. This is the most critical stage of the scheduling process. Continuing with the car analogy, you might score the three cars within your budget based on factors such as fuel efficiency, color, and safety ratings.

- **Picking**: In this final phase, the scheduler selects the worker with the highest or lowest score from the scoring phase.

**Manager**

The **Manager** consists of four main components: the **Scheduler**, **API**, **Task Store**, and **Workers**.

- The **Scheduler** used here is the one described above.

- The **API** is the primary interface for interacting with the system. Users can submit tasks, request task cancellations, and query the status of tasks and workers via the API.

- The **Task Store** is how the manager reliably tracks all tasks in the system. This tracking ensures sound scheduling decisions and enables the manager to provide accurate information about task and worker statuses to users.

- **Workers**: The manager oversees a collection of workers, monitoring their health and metrics in real time. Metrics include the number of tasks currently running on a worker, available memory, CPU load, and free disk space. These metrics, combined with data from the Task Store, support scheduling decisions.

**Worker**

The **Worker** consists of four main components: **API**, **Task Runtime**, **Task Storage**, and **Metrics**.

- Like the manager, the worker also has an **API** module, but its purpose differs. The worker's API primarily serves the manager. Through this API, the manager can send task assignment, cancellation, and retry requests to the worker, as well as retrieve the worker's status metrics.

- The **Task Runtime** is implemented using Docker (via the Docker Go client).

- The worker tracks the tasks it is responsible for using **Task Storage**.

- **Metrics**: The worker provides metrics about its current state, including task execution details and resource utilization, which it shares via its API for use by the manager.

> Development Environment: Ubuntu 22.04 arm64

## Start And API Test

1. start docker desktop Engine on your computer

2. `go build -o cube main.go`

3. `go install` (Make sure that $GOPATH/bin or $GOBIN has been added to your PATH environment variable)

4. start three worker using cube cli

```
cube worker --host 127.0.0.1 --port 5000 --dbtype "persistent"
cube worker --host 127.0.0.1 --port 5001 --dbtype "persistent"
cube worker --host 127.0.0.1 --port 5002 --dbtype "persistent"
```

5. start manager to receive task request and scedule to different worker

> cube manager -w 'localhost:5000,localhost:5001,localhost:5002'

6. simulate a user to send a task to the manager

```
cube run -m "localhost:5556" -f "task1.json"
cube run -m "localhost:5556" -f "task2.json"
cube run -m "localhost:5556" -f "task3.json"
cube run -m "localhost:5556" -f "task4.json"
cube run -m "localhost:5556" -f "task5.json"
```

7. check node state

```
$ cube node                                                                   
NAME               MEMORY (MiB)     DISK (GiB)     ROLE       TASKS     
localhost:5000     2014             66             worker     1         
localhost:5001     2014             66             worker     2         
localhost:5002     2014             66             worker     1 
// because three worker running on same machine
```

8. check task state

```
$ cube status
ID                                       NAME                 CREATED                    STATE       CONTAINERNAME        IMAGE                      
95fbe134-7f19-468a-acfc-c7753e5b4cd2     test-chapter-9.5     Less than a second ago     Running     test-chapter-9.5     sun4965485/echo-smy:v1     
bb1d59ef-9fc1-4e4b-a44d-db571eeed203     test-chapter-9.1     Less than a second ago     Running     test-chapter-9.1     sun4965485/echo-smy:v1     
21b23589-5d2d-4731-b5c9-a97e9832d021     test-chapter-9.2     Less than a second ago     Running     test-chapter-9.2     sun4965485/echo-smy:v1     
95fbe134-7f19-496a-acfc-c7853e5b4cd2     test-chapter-9.3     Less than a second ago     Running     test-chapter-9.3     sun4965485/echo-smy:v1     
95fbe134-7f19-496a-acfc-c7753e5b4cd2     test-chapter-9.4     Less than a second ago     Running     test-chapter-9.4     sun4965485/echo-smy:v1
```

9. stop task

```
cube stop 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/12 16:49:36 Task 21b23589-5d2d-4731-b5c9-a97e9832d021 has been stopped.

cube status
21b23589-5d2d-4731-b5c9-a97e9832d021     test-chapter-9.2     Less than a second ago     Completed     test-chapter-9.2     sun4965485/echo-smy:v1 
```

## Other ability

### Container port probe

```
2024/09/12 16:49:32 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7779"}}}
2024/09/12 16:49:32 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7802"}}}
2024/09/12 16:49:32 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7801"}}}
2024/09/12 16:49:32 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7800"}}}
2024/09/12 16:49:32 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7778"}}}
```

### Update task state

```
2024/09/12 16:46:27 Attempting to update task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/12 16:46:27 Attempting to update task 95fbe134-7f19-468a-acfc-c7753e5b4cd2
2024/09/12 16:46:27 Attempting to update task 95fbe134-7f19-496a-acfc-c7753e5b4cd2
2024/09/12 16:46:27 Attempting to update task 95fbe134-7f19-496a-acfc-c7853e5b4cd2
2024/09/12 16:46:27 Attempting to update task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
```

### Task scedule

```
manager log:
2024/09/12 17:29:19 Add event {a7aa1d44-08f6-443e-9378-f5884311018e 2 0001-01-01 00:00:00 +0000 UTC {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7778]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/12 17:29:19 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/12 17:29:27 [manager] selected worker localhost:5001 for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/12 17:29:27 [manager] received response from worker

worker log:
2024/09/12 17:29:27 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
7777/tcp:7778]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}
{"status":"Pulling from sun4965485/echo-smy","id":"v1"}
{"status":"Digest: sha256:b3a6951a31ab9ba821c95815ccc16de992fd00019fab37ed607514e61cf6f6fe"}
{"status":"Status: Image is up to date for sun4965485/echo-smy:v1"}
2024/09/12 17:29:30 task bb1d59ef-9fc1-4e4b-a44d-db571eeed203 Running on container 618b99d1610c99fbf24b115f41a2cd3a893323c2dff64c4c8eb7da294c1cd2b9
2024/09/12 17:29:43 updateTasks return resp Networking is nat.PortMap{"7777/tcp":[]nat.PortBinding{nat.PortBinding{HostIP:"127.0.0.1", HostPort:"7778"}}}
```

### One-click container cleaning     
```
$ cube prune                               
Stopping container d4ab751fdd1d...
2024/09/12 17:20:24 Container d4ab751fdd1d stopped: d4ab751fdd1d
2024/09/12 17:20:24 Container d4ab751fdd1d removed: d4ab751fdd1d

Stopping container 700cc7b217b4...
2024/09/12 17:20:24 Container 700cc7b217b4 stopped: 700cc7b217b4
2024/09/12 17:20:24 Container 700cc7b217b4 removed: 700cc7b217b4

Stopping container 32ec47f15d30...
2024/09/12 17:20:25 Container 32ec47f15d30 stopped: 32ec47f15d30
2024/09/12 17:20:25 Container 32ec47f15d30 removed: 32ec47f15d30

Stopping container 8948cacd7e29...
2024/09/12 17:20:25 Container 8948cacd7e29 stopped: 8948cacd7e29
2024/09/12 17:20:25 Container 8948cacd7e29 removed: 8948cacd7e29
```

### Collect cpu, disk and memory stat

```
2024/09/12 19:14:27 collect stats from worker http://localhost:5000 success, CPU detail: linux.CPUStat{Id:cpu, User:1031250, Nice:1541, System:399910, Idle:18235973, IOWait:14744, IRQ:0, SoftIRQ:11676, Steal:0, Guest:0, GuestNice:0}
2024/09/12 19:14:27 collect stats from worker http://localhost:5000 success, Disk detail: linux.Disk{All:66205626368, Used:17241092096, Free:48964534272, FreeInodes:3702938}
2024/09/12 19:14:27 collect stats from worker http://localhost:5000 success, Memory detail: linux.MemInfo{MemTotal:2014312, MemFree:92044, MemAvailable:545224, Buffers:25916, Cached:649952, SwapCached:54664, Active:791872, Inactive:635028, ActiveAnon:455416, InactiveAnon:525576, ActiveFile:336456, InactiveFile:109452, Unevictable:223784, Mlocked:80, SwapTotal:2097148, SwapFree:1155720, Dirty:244, Writeback:0, AnonPages:961424, Mapped:200268, Shmem:229960, Slab:193996, SReclaimable:97384, SUnreclaim:96612, KernelStack:11248, PageTables:28808, NFS_Unstable:0, Bounce:0, WritebackTmp:0, CommitLimit:3104304, Committed_AS:7954852, VmallocTotal:133143592960, VmallocUsed:29556, VmallocChunk:0, HardwareCorrupted:0, AnonHugePages:0, HugePages_Total:0, HugePages_Free:0, HugePages_Rsvd:0, HugePages_Surp:0, Hugepagesize:2048, DirectMap4k:0, DirectMap2M:0, DirectMap1G:0}
```

### Health probe

```
2024/09/12 19:16:27 Performing task health check
2024/09/12 19:16:27 url is : http://localhost:5000/tasks/
2024/09/12 19:16:27 url is : http://localhost:5001/tasks/
2024/09/12 19:16:27 url is : http://localhost:5002/tasks/
2024/09/12 19:16:27 Task health checks completed

2024/09/12 19:35:17 collect stats from worker http://localhost:5002 success, Memory detail:
linux.MemInfo{MemTotal:2014312, MemFree:120320, MemAvailable:497580, Buffers:53984, Cached:487972,
SwapCached:66796, Active:638788, Inactive:672392, ActiveAnon:451248, InactiveAnon:551736, ActiveFile:187540,
InactiveFile:120656, Unevictable:229476, Mlocked:80, SwapTotal:2097148, SwapFree:1151276, Dirty:72, Writeback:0,
AnonPages:983380, Mapped:173624, Shmem:233760, Slab:256040, SReclaimable:159176, SUnreclaim:96864, KernelStack:11232,
PageTables:28980, NFS_Unstable:0, Bounce:0, WritebackTmp:0, CommitLimit:3104304, Committed_AS:8054100, VmallocTotal:133143592960,
VmallocUsed:29572, VmallocChunk:0, HardwareCorrupted:0, AnonHugePages:0, HugePages_Total:0, HugePages_Free:0, HugePages_Rsvd:0, HugePages_Surp:0, Hugepagesize:2048, DirectMap4k:0, DirectMap2M:0, DirectMap1G:0}


2024/09/12 19:36:17 collect stats from worker http://localhost:5002 success, Memory detail:
linux.MemInfo{MemTotal:2014312, MemFree:93096, MemAvailable:472472, Buffers:54108, Cached:486956,
SwapCached:67052, Active:640584, Inactive:718120, ActiveAnon:451456, InactiveAnon:596968, ActiveFile:189128,
InactiveFile:121152, Unevictable:216152, Mlocked:80, SwapTotal:2097148, SwapFree:1151788, Dirty:244, Writeback:0,
AnonPages:1018376, Mapped:175420, Shmem:230784, Slab:256064, SReclaimable:159208, SUnreclaim:96856, KernelStack:11248,
PageTables:28980, NFS_Unstable:0, Bounce:0, WritebackTmp:0, CommitLimit:3104304, Committed_AS:8061180, VmallocTotal:133143592960,
VmallocUsed:29604, VmallocChunk:0, HardwareCorrupted:0, AnonHugePages:0, HugePages_Total:0, HugePages_Free:0, HugePages_Rsvd:0, HugePages_Surp:0, Hugepagesize:2048, DirectMap4k:0, DirectMap2M:0, DirectMap1G:0}
```
