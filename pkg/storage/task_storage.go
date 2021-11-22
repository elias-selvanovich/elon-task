package persistence

import (
	"elon-task/pkg/model"
	"log"
)

var tasks map[string]model.Task

type TaskPersistence interface {
	SaveTask(model.Task) (bool, error)
	GetTask(id uint) (model.Task, error)
}

type taskStorage struct {
	logger log.Logger
}


