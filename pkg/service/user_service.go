package service

import (
	"elon-task/pkg/model"
	"elon-task/pkg/storage"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUsers() []*model.User
	GetUser(email string) *model.User
}

type userService struct {
	userStorage storage.UserStorage
}

func NewUserService(u storage.UserStorage) *userService {
	return &userService{u}
}

func (s *userService) GetUser(email string) *model.User {
	return s.userStorage.GetUser(email)
}

func (s *userService) GetUsers() []*model.User {
	return s.userStorage.GetUsers()
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userStorage.SaveUser(user)
}
