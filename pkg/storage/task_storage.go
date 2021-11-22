package storage

import (
	"elon-task/pkg/model"
	"errors"
	"fmt"
	"log"
	"time"
)

var currIndex = uint(0)
var tasks = map[string][]*model.Task{}

type TaskStorage interface {
	GetTasks(email string) ([]*model.Task, error)
	UpdateTask(task *model.Task) error
	SaveTask(task *model.Task) error
	DeleteTask(task *model.Task) error
}

type taskStorage struct{}

func NewTaskStorage() *taskStorage {
	return &taskStorage{}
}

func (s *taskStorage) GetTasks(email string) ([]*model.Task, error) {
	if t, ok := tasks[email]; ok {
		return t, nil
	} else {
		log.Println(fmt.Sprintf("Tasks not found for user %s", email))
		return []*model.Task{}, nil
	}
}

func (s *taskStorage) UpdateTask(task *model.Task) error {
	if t, ok := tasks[task.Email]; ok {
		// loop by index
		for i := range t {
			if t[i].Id == task.Id {
				task.CreatedTime = t[i].CreatedTime
				task.UpdatedTime = time.Now()
				t[i] = task
				return nil
			}
		}
	}

	return errors.New("Task not found")
}

func (s *taskStorage) SaveTask(task *model.Task) error {
	currIndex += 1

	task.Id = currIndex
	task.CreatedTime = time.Now()
	// if email not found then create new slice
	if _, ok := tasks[task.Email]; !ok {
		tasks[task.Email] = make([]*model.Task, 0)
	}

	tasksArr := tasks[task.Email]

	tasksArr = append(tasksArr, task)
	tasks[task.Email] = tasksArr

	return nil
}

func (s *taskStorage) DeleteTask(task *model.Task) error {
	if t, ok := tasks[task.Email]; ok {
		// loop by index
		for i := range t {
			if t[i].Id == task.Id {
				task.CreatedTime = t[i].CreatedTime
				task.UpdatedTime = t[i].UpdatedTime
				task.RemovedTime = time.Now()
				t[i] = task
				return nil
			}
		}
	}

	return errors.New("Task not found")
}
