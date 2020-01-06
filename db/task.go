package db

import (
	"encoding/binary"
	"time"

	bolt "go.etcd.io/bbolt"
)

type App struct {
	DB *bolt.DB
}

func (a *App) Init(dbPath string) error {
	var err error
	a.DB, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: time.Second})
	if err != nil {
		return err
	}
	return a.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		return err
	})
}

func (a *App) Create(task string) (int, error) {
	var id int
	err := a.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return id, nil
}

type Task struct {
	Key int
	Val string
}

func (a *App) AllTasks() ([]Task, error) {
	var tasks []Task
	err := a.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key: btoi(k),
				Val: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (a *App) DeleteTask(key int) error {
	return a.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Delete(itob(key))
	})
}

func (a *App) Close() error {
	return a.DB.Close()
}

func itob(n int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

func btoi(b []byte) int {
	int64 := binary.BigEndian.Uint64(b)
	return int(int64)
}
