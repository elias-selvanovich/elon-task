package model

import "time"

const (
	TASK_CREATED     = iota
	TASK_IN_PROGRESS = 1
	TASK_FINISHED    = 2
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Task struct {
	Id          uint      `json:"id"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	Status      uint      `json:"status"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	RemovedTime time.Time `json:"removed_time"`
}
