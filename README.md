# cube-orchestrator

## Start And API Test

> start docker desktop Engine on your computer

1. export CUBE_HOST=localhost
2. export CUBE_PORT=5555
3. go run main.go
4. curl -v localhost:5555/tasks | jq
5. 

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
    "ID": "266592cd-960d-4091-981c-8c25c44b1018",
    "State": 2,
    "Task": {
        "State": 1,
        "ID": "266592cd-960d-4091-981c-8c25c44b1018",
        "Name": "test-chapter-5-1",
        "Image": "strm/helloworld-http"
    }
}' \
http://127.0.0.1:5555/tasks


or


curl -v --request POST \                                                            
--header 'Content-Type: application/json' \
--data @post_task.json \
http://127.0.0.1:5555/tasks

6. curl -v localhost:5555/tasks | jq
7. docker ps --format "table {{.ID}}\t{{.Image}}\t{{.Status}}\t{{.Names}}"
8. curl -v --request DELETE "localhost:5555/tasks/266592cd-960d-4091-981c-8c25c44b1018"
9. curl -v localhost:5555/tasks | jq
10. check if the task: 266592cd-960d-4091-981c-8c25c44b1018 state is 3




## State-Machine Test

- curl -v http://127.0.0.1:5555/tasks | jq
- simulate manager distribute a task to worker

curl -v --request POST \                
--header 'Content-Type: application/json' \
--data @add_task.json \
http://127.0.0.1:5555/tasks

- task execution completed

curl -v --request POST \                
--header 'Content-Type: application/json' \
--data @stop_task.json \
http://127.0.0.1:5555/tasks