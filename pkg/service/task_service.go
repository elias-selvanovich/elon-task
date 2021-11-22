package service

import (
	"elon-task/pkg/model"
	"elon-task/pkg/storage"
)

type TaskService interface {
	GetTasks(email string) ([]*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(task *model.Task) error
}

type taskService struct {
	taskStorage storage.TaskStorage
}

func NewTaskService(t storage.TaskStorage) *taskService {
	return &taskService{t}
}

func (s *taskService) GetTasks(email string) ([]*model.Task, error) {
	return s.taskStorage.GetTasks(email)
}

func (s *taskService) UpdateTask(task *model.Task) error {
	return s.taskStorage.UpdateTask(task)
}

func (s *taskService) CreateTask(task *model.Task) error {
	return s.taskStorage.SaveTask(task)
}

func (s *taskService) DeleteTask(task *model.Task) error {
	return s.taskStorage.DeleteTask(task)
}
