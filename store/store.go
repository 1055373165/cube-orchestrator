package store

import (
	"cube/task"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

type Store interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{}) error
	List() (interface{}, error)
	Count() (int, error)
}

type InMemoryTaskStore struct {
	Db map[string]*task.Task
}

func NewInMemoryTaskStore() *InMemoryTaskStore {
	return &InMemoryTaskStore{
		Db: make(map[string]*task.Task),
	}
}

func (i *InMemoryTaskStore) Put(key string, value interface{}) error {
	t, ok := value.(*task.Task)
	if !ok {
		return fmt.Errorf("value %v is not a task.Task type", value)
	}
	i.Db[key] = t
	return nil
}

func (i *InMemoryTaskStore) Get(key string) (interface{}, error) {
	t, ok := i.Db[key]
	if !ok {
		return nil, fmt.Errorf("task with key %s does not exist", key)
	}
	return t, nil
}

func (i *InMemoryTaskStore) List() (interface{}, error) {
	var tasks []*task.Task
	for _, t := range i.Db {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (i *InMemoryTaskStore) Count() (int, error) {
	return len(i.Db), nil
}

type InMemoryTaskEventStore struct {
	Db map[string]*task.TaskEvent
}

func NewInMemoryTaskEventStore() *InMemoryTaskEventStore {
	return &InMemoryTaskEventStore{
		Db: make(map[string]*task.TaskEvent),
	}
}

func (i *InMemoryTaskEventStore) Put(key string, value interface{}) error {
	t, ok := value.(*task.TaskEvent)
	if !ok {
		return fmt.Errorf("value %v is not a task.TaskEvent type", value)
	}
	i.Db[key] = t
	return nil
}

func (i *InMemoryTaskEventStore) Get(key string) (interface{}, error) {
	t, ok := i.Db[key]
	if !ok {
		return nil, fmt.Errorf("task event with key %s does not exist", key)
	}
	return t, nil
}

func (i *InMemoryTaskEventStore) List() (interface{}, error) {
	var events []*task.TaskEvent
	for _, t := range i.Db {
		events = append(events, t)
	}
	return events, nil
}

func (i *InMemoryTaskEventStore) Count() (int, error) {
	return len(i.Db), nil
}

type TaskStore struct {
	DbFile   string
	FileMode os.FileMode
	Db       *bolt.DB
	Bucket   string
}

func NewTaskStore(filename string, mode os.FileMode, bucket string) (*TaskStore, error) {
	db, err := bolt.Open(filename, mode, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v", filename)
	}

	t := &TaskStore{
		DbFile:   filename,
		FileMode: mode,
		Bucket:   bucket,
		Db:       db,
	}

	err = t.CreateBucket()
	if err != nil {
		log.Printf("bucket already exists, will use it instead of creating new one")
	}

	return t, nil
}

func (ts *TaskStore) CreateBucket() error {
	err := ts.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(ts.Bucket))
		if err != nil {
			return fmt.Errorf("create bucket %s: %s", ts.Bucket, err)
		}
		return nil
	})
	return err
}

func (ts *TaskStore) Close() {
	ts.Db.Close()
}

func (ts *TaskStore) Get(key string) (interface{}, error) {
	var task task.Task
	err := ts.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ts.Bucket))
		t := b.Get([]byte(key))
		if t == nil {
			return fmt.Errorf("%s not found", key)
		}
		err := json.Unmarshal(t, &task)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (ts *TaskStore) Put(key string, value interface{}) error {
	err := ts.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ts.Bucket))
		t, err := json.Marshal(value.(*task.Task))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), t)
		if err != nil {
			return fmt.Errorf("unable to save %s item", key)
		}
		return nil
	})
	return err
}

func (ts *TaskStore) List() (interface{}, error) {
	var tasks []*task.Task
	err := ts.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ts.Bucket))
		err := b.ForEach(func(k, v []byte) error {
			var task task.Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			tasks = append(tasks, &task)
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskStore) Count() (int, error) {
	var taskCount int
	err := ts.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ts.Bucket))
		err := b.ForEach(func(k, v []byte) error {
			taskCount++
			return nil
		})
		return err
	})
	if err != nil {
		return -1, err
	}
	return taskCount, nil
}

type TaskEventStore struct {
	DbFile   string
	FileMode os.FileMode
	Db       *bolt.DB
	Bucket   string
}

func NewTaskEventStore(file string, mode os.FileMode, bucket string) (*TaskEventStore, error) {
	db, err := bolt.Open(file, mode, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v", file)
	}
	e := TaskEventStore{
		DbFile:   file,
		FileMode: mode,
		Db:       db,
		Bucket:   bucket,
	}

	err = e.CreateBucket()
	if err != nil {
		log.Printf("bucket already exists, will use it instead of creating new one")
	}

	return &e, nil
}

func (tes *TaskEventStore) Close() {
	tes.Db.Close()
}

func (tes *TaskEventStore) CreateBucket() error {
	return tes.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(tes.Bucket))
		if err != nil {
			log.Printf("bucket already exists, will use it instead of creating new one")
			return err
		}
		return nil
	})
}

func (e *TaskEventStore) Get(key string) (interface{}, error) {
	var event task.TaskEvent
	err := e.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(e.Bucket))
		v := b.Get([]byte(key))
		if v == nil {
			return fmt.Errorf("event %s not found", key)
		}
		err := json.Unmarshal(v, &event)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *TaskEventStore) Put(key string, value interface{}) error {
	return e.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(e.Bucket))

		buf, err := json.Marshal(value.(task.TaskEvent))
		if err != nil {
			return err
		}

		err = b.Put([]byte(key), buf)
		if err != nil {
			log.Printf("unable to save item %s", key)
			return err
		}
		return nil
	})
}

func (e *TaskEventStore) List() (interface{}, error) {
	var events []*task.TaskEvent
	err := e.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(e.Bucket))
		b.ForEach(func(k, v []byte) error {
			var te task.TaskEvent
			err := json.Unmarshal(v, &te)
			if err != nil {
				return err
			}
			events = append(events, &te)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (e *TaskEventStore) Count() (int, error) {
	eventCount := 0
	err := e.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(e.Bucket))
		b.ForEach(func(k, v []byte) error {
			eventCount++
			return nil
		})
		return nil
	})
	if err != nil {
		return -1, err
	}
	return eventCount, nil
}
