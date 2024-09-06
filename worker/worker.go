package worker

import (
	"cube/task"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// 1. Run tasks as Docker containers
// 2. Accept tasks to run from a manager
// 3. Provide relevant statistics to the manager  for the purpost of scheduling tasks
// 4. Keep track of its tasks and their state
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

// 1. Create an instance of the Docker struct that allows us to talk to the Docker daemon using the Docker SDK.
// 2. Call the Stop() method on the Docker struct.
// 3. Check whether there were any errors in stopping the task.
// 4. Update the FinishTime field on the task t.
// 5. Save the updated task t to the worker's DB field
// 6. Print hte informative message and return the result of the operation.
func (w *Worker) StopTask(t task.Task) task.DockerResult {
	config := task.NewConfig(&t)
	d := task.NewDocker(config)

	result := d.Stop(t.ContainerID)
	if result.Error != nil {
		log.Printf("Error stopping container %v: %v\n", t.ContainerID, result.Error)
		return result
	}

	t.FinishTime = time.Now().UTC()
	t.State = task.Completed
	w.Db[t.ID] = &t // Simulate database updates.

	log.Printf("Stopped and removed container %v for task %v\n", t.ContainerID, t.ID)
	return result
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

// Guaranteeing idempotency of multiple-run status.
// 1. Pull a task off the queue.
// 2. Convert it from an interface to a task.Task type.
// 3. Retrieve the task from the worker's Db.
// 4. Check whether the state transition is valid.
// 5. If the task from the queue is in the state Scheduled, call StartTask.
// 6. If the task from the queue is in the state Completed, call StopTask.
// 6. Else, there is an invalid transition, so return an error.
func (w *Worker) RunTask() task.DockerResult {
	t := w.Queue.Dequeue()
	if t == nil {
		log.Println("No tasks in the queue")
		return task.DockerResult{Error: nil}
	}

	taskQueued := t.(task.Task)
	taskPersisted := w.Db[taskQueued.ID]
	if taskPersisted == nil {
		taskPersisted = &taskQueued
		w.Db[taskQueued.ID] = &taskQueued
	}

	var result task.DockerResult
	if task.ValidStateTransition(taskPersisted.State, taskQueued.State) {
		switch taskQueued.State {
		case task.Scheduled:
			result = w.StartTask(taskQueued)
		case task.Completed:
			result = w.StopTask(taskQueued)
		default:
			result.Error = errors.New("we should not get here")
		}
	} else {
		err := fmt.Errorf("invalid transition from %v to %v",
			taskPersisted.State, taskQueued.State)
		result.Error = err
	}
	return result
}

func (w *Worker) StartTask(t task.Task) task.DockerResult {
	c := task.NewConfig(&t)
	d := task.NewDocker(c)

	result := d.Run()
	if result.Error != nil {
		log.Printf("Error running task %v: %v\n", t.ID, result.Error)
		t.State = task.Failed
		w.Db[t.ID] = &t
		return result
	}

	t.ContainerID = result.ContainerId
	t.State = task.Running
	w.Db[t.ID] = &t

	log.Printf("task %v Running on container %v\n", t.ID, t.ContainerID)
	return result
}

func (w *Worker) AddTask(t task.Task) {
	w.Queue.Enqueue(t)
}
