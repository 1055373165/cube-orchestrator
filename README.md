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