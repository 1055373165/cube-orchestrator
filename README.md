# cube-orchestrator

## Start And API Test

> start docker desktop Engine on your computer

### Service start

docker run -d -p 7777:7777 --name echo sun4965485/echo-smy:v1

```
CUBE_WORKER_HOST=127.0.0.1 \
CUBE_WORKER_PORT=5000 \
CUBE_MANAGER_HOST=127.0.0.1 \
CUBE_MANAGER_PORT=5556 \
go run main.go
```
### POST new task

curl -v --request POST --header 'Content-Type: application/json' --data @task1.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @task2.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @task3.json localhost:5556/tasks
curl -v --request POST --header 'Content-Type: application/json' --data @task4.json localhost:5556/tasks


2024/09/08 20:35:07 Pulled {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queue
2024/09/08 20:35:07 [manager] selected worker 127.0.0.1:5001 for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203


2024/09/08 20:38:07 Pulled {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {21b23589-5d2d-4731-b5c9-a97e9832d021  test-chapter-9.2 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queue
2024/09/08 20:38:07 [manager] selected worker 127.0.0.1:5002 for task 21b23589-5d2d-4731-b5c9-a97e9832d021


2024/09/08 20:38:17 Pulled {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {95fbe134-7f19-496a-acfc-c7853e5b4cd2  test-chapter-9.3 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queue
2024/09/08 20:38:17 [manager] selected worker 127.0.0.1:5000 for task 95fbe134-7f19-496a-acfc-c7853e5b4cd2

2024/09/08 20:39:57 Pulled {a7aa1d44-08f6-443e-9378-f5884313419e 2 0001-01-01 00:00:00 +0000 UTC {95fbe134-7f19-496a-acfc-c7753e5b4cd2  test-chapter-9.4 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queue
2024/09/08 20:39:57 [manager] selected worker 127.0.0.1:5001 for task 95fbe134-7f19-496a-acfc-c7753e5b4cd2


2024/09/08 22:48:37 Add event {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/08 22:48:37 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/08 22:48:37 Add event {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {21b23589-5d2d-4731-b5c9-a97e9832d021  test-chapter-9.2 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/08 22:48:37 Added task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/08 22:48:37 Add event {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {95fbe134-7f19-496a-acfc-c7853e5b4cd2  test-chapter-9.3 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/08 22:48:37 Added task 95fbe134-7f19-496a-acfc-c7853e5b4cd2
2024/09/08 22:48:37 Add event {a7aa1d44-08f6-443e-9378-f5884313419e 2 0001-01-01 00:00:00 +0000 UTC {95fbe134-7f19-496a-acfc-c7753e5b4cd2  test-chapter-9.4 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} to pending queue
2024/09/08 22:48:37 Added task 95fbe134-7f19-496a-acfc-c7753e5b4cd2


2024/09/08 22:48:44 Pulled {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {bb1d59ef-9fc1-4e4b-a44d-db571eeed203  test-chapter-9.1 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queue
2024/09/08 22:48:53 [manager] selected worker 127.0.0.1:5000 for task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/08 22:48:53 Added task bb1d59ef-9fc1-4e4b-a44d-db571eeed203
2024/09/08 22:48:53 [manager] received response from worker: task.Task{ID:uuid.UUID{0xbb, 0x1d, 0x59, 0xef, 0x9f, 0xc1, 0x4e, 0x4b, 0xa4, 0x4d, 0xdb, 0x57, 0x1e, 0xee, 0xd2, 0x3}, ContainerID:"", Name:"test-chapter-9.1", State:1, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet{"7777/tcp":struct {}{}}, HostPorts:nat.PortMap(nil), PortBindings:map[string]string{"7777/tcp":"7777"}, RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"/health", RestartCount:0}
{"status":"Pulling from sun4965485/echo-smy","id":"v1"}
{"status":"Digest: sha256:b3a6951a31ab9ba821c95815ccc16de992fd00019fab37ed607514e61cf6f6fe"}
{"status":"Status: Image is up to date for sun4965485/echo-smy:v1"}
2024/09/08 14:48:57 Listening on http://localhost:7777
2024/09/08 22:48:57 task bb1d59ef-9fc1-4e4b-a44d-db571eeed203 Running on container 1e9231d0d90d7c3b5eefd2721a28506a66af5fbfe310629e45787b8550d58a77



2024/09/08 22:49:03 Pulled {a7aa1d44-08f6-443e-9378-f5884311019e 2 0001-01-01 00:00:00 +0000 UTC {21b23589-5d2d-4731-b5c9-a97e9832d021  test-chapter-9.2 1 sun4965485/echo-smy:v1 0 0 0 map[7777/tcp:{}] map[] map[7777/tcp:7777]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC /health 0}} off pending queue
2024/09/08 22:49:12 [manager] selected worker 127.0.0.1:5001 for task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/08 22:49:12 Added task 21b23589-5d2d-4731-b5c9-a97e9832d021
2024/09/08 22:49:12 [manager] received response from worker: task.Task{ID:uuid.UUID{0x21, 0xb2, 0x35, 0x89, 0x5d, 0x2d, 0x47, 0x31, 0xb5, 0xc9, 0xa9, 0x7e, 0x98, 0x32, 0xd0, 0x21}, ContainerID:"", Name:"test-chapter-9.2", State:1, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet{"7777/tcp":struct {}{}}, HostPorts:nat.PortMap(nil), PortBindings:map[string]string{"7777/tcp":"7777"}, RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"/health", RestartCount:0}
{"status":"Pulling from sun4965485/echo-smy","id":"v1"}
{"status":"Digest: sha256:b3a6951a31ab9ba821c95815ccc16de992fd00019fab37ed607514e61cf6f6fe"}
{"status":"Status: Image is up to date for sun4965485/echo-smy:v1"}
2024/09/08 22:49:17 No tasks to process currently.
2024/09/08 22:49:17 Sleeping 10 time seconds
2024/09/08 14:49:17 Listening on http://localhost:7777
2024/09/08 22:49:17 task 21b23589-5d2d-4731-b5c9-a97e9832d021 Running on container d36869ab5de030efcc688ffcf8b3fcef78a13037b8efc3bb8c46a83184a983a6

2024/09/08 22:49:31 [manager] selected worker 127.0.0.1:5002 for task 95fbe134-7f19-496a-acfc-c7853e5b4cd2
2024/09/08 22:49:31 Added task 95fbe134-7f19-496a-acfc-c7853e5b4cd2
2024/09/08 22:49:31 [manager] received response from worker: task.Task{ID:uuid.UUID{0x95, 0xfb, 0xe1, 0x34, 0x7f, 0x19, 0x49, 0x6a, 0xac, 0xfc, 0xc7, 0x85, 0x3e, 0x5b, 0x4c, 0xd2}, ContainerID:"", Name:"test-chapter-9.3", State:1, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet{"7777/tcp":struct {}{}}, HostPorts:nat.PortMap(nil), PortBindings:map[string]string{"7777/tcp":"7777"}, RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"/health", RestartCount:0}

2024/09/08 22:49:50 [manager] selected worker 127.0.0.1:5000 for task 95fbe134-7f19-496a-acfc-c7753e5b4cd2
2024/09/08 22:49:50 Added task 95fbe134-7f19-496a-acfc-c7753e5b4cd2
2024/09/08 22:49:50 [manager] received response from worker: task.Task{ID:uuid.UUID{0x95, 0xfb, 0xe1, 0x34, 0x7f, 0x19, 0x49, 0x6a, 0xac, 0xfc, 0xc7, 0x75, 0x3e, 0x5b, 0x4c, 0xd2}, ContainerID:"", Name:"test-chapter-9.4", State:1, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet{"7777/tcp":struct {}{}}, HostPorts:nat.PortMap(nil), PortBindings:map[string]string{"7777/tcp":"7777"}, RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"/health", RestartCount:0}


2024/09/08 22:52:29 [manager] selected worker 127.0.0.1:5001 for task 95fbe134-7f19-468a-acfc-c7753e5b4cd2
2024/09/08 22:52:29 Added task 95fbe134-7f19-468a-acfc-c7753e5b4cd2
2024/09/08 22:52:29 [manager] received response from worker: task.Task{ID:uuid.UUID{0x95, 0xfb, 0xe1, 0x34, 0x7f, 0x19, 0x46, 0x8a, 0xac, 0xfc, 0xc7, 0x75, 0x3e, 0x5b, 0x4c, 0xd2}, ContainerID:"", Name:"test-chapter-9.5", State:1, Image:"sun4965485/echo-smy:v1", CPU:0, Memory:0, Disk:0, ExposedPorts:nat.PortSet{"7777/tcp":struct {}{}}, HostPorts:nat.PortMap(nil), PortBindings:map[string]string{"7777/tcp":"7777"}, RestartPolicy:"", StartTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), FinishTime:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), HealthCheck:"/health", RestartCount:0}
2024/09/08 22:52:29 Sleeping for 10 seconds