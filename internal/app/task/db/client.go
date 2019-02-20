package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// TaskKey defines all the keys that are available for a task in DB.
type TaskKey string

// Key Constants represent the keys of the fields stored in DB.
const (
	IDKey     = "id"
	DataKey   = "data"
	StatusKey = "status"
)

// TaskStatus defines the DB status of a task.
type TaskStatus string

// Status Constants represent the possible status of the tasks stored in DB.
const (
	PendingStatus   TaskStatus = "P"
	CompletedStatus TaskStatus = "C"
)

func (status TaskStatus) String() string {
	return string(status)
}

// Task defines the DB representation of a task.
type Task struct {
	ID     int
	Data   string
	Status TaskStatus
}

var client *redis.Client
var id = 0

const queryPageSize int64 = 100

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	if _, err := client.Ping().Result(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("redis connection established\n")
}

// Insert adds the data to the existing records in the database.
// If the insertion is unsuccessful, it returns an error.
func Insert(data string) error {
	id, err := nextID()
	if err != nil {
		return err
	}
	if cmd := client.HSet(id, DataKey, data); cmd.Err() != nil {
		return cmd.Err()
	}
	if cmd := client.HSet(id, StatusKey, PendingStatus.String()); cmd.Err() != nil {
		idx, _ := strconv.Atoi(id)
		_ = Delete(idx)
		return cmd.Err()
	}
	return nil
}

// UpdateStatus updates the status of the task with the given ID.
func UpdateStatus(id int, status string) error {
	return client.HSet(strconv.Itoa(id), StatusKey, status).Err()
}

// FetchAll returns all the tasks present in the database.
func FetchAll() ([]Task, error) {
	var cursor uint64
	var tasks []Task
	for {
		keys, cursor, err := client.Scan(cursor, "*", queryPageSize).Result()
		if err != nil {
			return nil, err
		}
		for _, key := range keys {
			if key == IDKey {
				continue
			}
			cmd := client.HGetAll(key)
			res, err := cmd.Result()
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, buildTask(key, res))
		}
		if cursor == 0 {
			break
		}
	}
	return tasks, nil
}

// Delete deletes the data present for the given id.
// If the id is not present, it returns an error.
func Delete(id int) error {
	return client.Del(strconv.Itoa(id)).Err()
}

func nextID() (string, error) {
	cmd := client.Incr(IDKey)
	id, err := cmd.Result()
	return strconv.Itoa(int(id)), err
}

func buildTask(id string, res map[string]string) Task {
	idx, _ := strconv.Atoi(id)
	return Task{
		ID:     idx,
		Data:   res[DataKey],
		Status: TaskStatus(res[StatusKey]),
	}
}
