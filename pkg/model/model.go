package model

import "time"

type User struct {
	Id    uint
	Name  string
	Email string
}

type Task struct {
	Id          uint      `json:"id"`
	UserId      uint      `json:"user_id"`
	Description string    `json:"description"`
	Status      uint      `json:"status"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	RemovedTime time.Time `json:"removed_time"`
}
