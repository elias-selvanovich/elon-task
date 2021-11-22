package storage

import (
	"elon-task/pkg/model"
	"fmt"
	"log"
)

var users = make(map[string]*model.User, 0)

type UserStorage interface {
	GetUser(email string) *model.User
	GetUsers() []*model.User
	SaveUser(user *model.User) error
}

type userStorage struct{}

func NewUserStorage() *userStorage {
	return &userStorage{}
}

func (s *userStorage) GetUser(email string) *model.User {
	if u, ok := users[email]; ok {
		return u
	} else {
		log.Println(fmt.Sprintf("User %s not found", email))
		return nil
	}
}

func (s *userStorage) GetUsers() []*model.User {
	userArr := make([]*model.User, 0)
	for _, u := range users {
		userArr = append(userArr, u)
	}

	return userArr
}

func (s *userStorage) SaveUser(user *model.User) error {
	if _, ok := users[user.Email]; !ok {
		users[user.Email] = user
	} else {
		log.Println(fmt.Sprintf("User %s already stored", user.Email))
	}

	return nil
}
