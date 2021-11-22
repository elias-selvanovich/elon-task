package service

import (
	"elon-task/pkg/model"
	"log"
)

type UserService interface {
	CreateUser(name, email string) model.User
	GetUsers() []model.User
	GetUser(id uint) model.User
}

type userService struct {
	logger log.Logger
}
