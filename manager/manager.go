package manager

import (
	"bytes"
	"cube/task"
	"cube/worker"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// 1. Accept requests from users to start or stop tasks
// 2. Schedule tasks onto worker machines
// 3. Keep track of tasks, their states, and the machine on which they run
type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string]*task.Task
	EventDb       map[string]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
	LastWorker    int
}

func New(workers []string) *Manager {
	taskDb := make(map[string]*task.Task)         // task uuid -> task object
	eventDb := make(map[string]*task.TaskEvent)   // event uuid -> task event object
	workerTaskMap := make(map[string][]uuid.UUID) // ip:port -> []taskUUID
	taskWorkerMap := make(map[uuid.UUID]string)   // task uuid -> ip:port
	for _, worker := range workers {
		workerTaskMap[worker] = []uuid.UUID{}
	}
	return &Manager{
		Pending:       *queue.New(),
		TaskDb:        taskDb,
		EventDb:       eventDb,
		Workers:       workers,
		WorkerTaskMap: workerTaskMap,
		TaskWorkerMap: taskWorkerMap,
		LastWorker:    0,
	}
}

func (m *Manager) SelectWorker() string {
	var newWorker int
	if m.LastWorker+1 < len(m.Workers) {
		newWorker = m.LastWorker + 1
		m.LastWorker++
	} else {
		newWorker = 0
		m.LastWorker = 0
	}

	return m.Workers[newWorker]
}

func (m *Manager) updateTasks() {
	for _, worker := range m.Workers {
		log.Printf("Checking worker %v for task updates\n", worker)
		url := fmt.Sprintf("http://%s/tasks", worker)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error connecting to %v: %v\n", worker, err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error sending request: %v\n", err.Error())
			return
		}

		var tasks []task.Task
		err = json.NewDecoder(resp.Body).Decode(&tasks)
		if err != nil {
			log.Printf("Error unmarshalling tasks: %v\n", err.Error())
			return
		}

		for _, task := range tasks {
			log.Printf("Attempting to update task %v\n", task.ID.String())

			_, ok := m.TaskDb[task.ID.String()]
			if !ok {
				log.Printf("Task with ID %s not found\n", task.ID.String())
				continue
			}

			if m.TaskDb[task.ID.String()].State != task.State {
				m.TaskDb[task.ID.String()].State = task.State
			}

			m.TaskDb[task.ID.String()].StartTime = task.StartTime
			m.TaskDb[task.ID.String()].FinishTime = task.FinishTime
			m.TaskDb[task.ID.String()].ContainerID = task.ContainerID // failover
		}
	}
}

func (m *Manager) UpdateTasks() {
	for {
		log.Println("Checking for task updates from workers")
		m.updateTasks()
		log.Println("Task updates completed")
		log.Println("Sleeping for 15 seconds")
		time.Sleep(15 * time.Second)
	}
}

func (m *Manager) ProcessTasks() {
	for {
		log.Println("Processing any tasks in the queue")
		m.SendWork()
		log.Println("Sleeping for 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func (m *Manager) SendWork() {
	if m.Pending.Len() == 0 {
		log.Println("No task event in the queue")
		return
	}

	w := m.SelectWorker()
	e := m.Pending.Dequeue()
	te := e.(task.TaskEvent)
	t := te.Task
	log.Printf("Pulled %v off pending queue\n", t)

	m.EventDb[te.ID.String()] = &te
	m.TaskDb[t.ID.String()] = &t
	m.WorkerTaskMap[w] = append(m.WorkerTaskMap[w], t.ID)
	m.TaskWorkerMap[t.ID] = w
	t.State = task.Scheduled

	data, err := json.Marshal(te)
	if err != nil {
		log.Printf("Unable to marshal task event object: %v.\n", t)
		return
	}

	url := fmt.Sprintf("http://%s/tasks", w)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error connecting to %v: %v\n", w, err.Error())
		m.Pending.Enqueue(te) // scheduled failed, re-enter queue
		return
	}

	d := json.NewDecoder(resp.Body)
	// if request failed, resp body is ErrorResponse
	if resp.StatusCode != http.StatusCreated {
		e := worker.ErrorResponse{}
		err := d.Decode(&e)
		if err != nil {
			fmt.Printf("Error decoding response: %s\n", err.Error())
			return
		}
		log.Printf("Response error (%d): %s\n", e.HTTPStateCode, e.Message)
		return
	}

	// if request sucess, resp body is task object
	t = task.Task{}
	err = d.Decode(&t)
	if err != nil {
		fmt.Printf("Error decoding response: %s\n", err.Error())
		return
	}

	log.Printf("%#v\n", t)
}

func (m *Manager) GetTasks() []*task.Task {
	tasks := []*task.Task{}
	for _, t := range m.TaskDb {
		tasks = append(tasks, t)
	}
	return tasks
}

func (m *Manager) AddTask(te task.TaskEvent) {
	m.Pending.Enqueue(te)
}
