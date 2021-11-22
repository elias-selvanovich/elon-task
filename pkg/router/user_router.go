package router

import (
	"elon-task/pkg/model"
	"elon-task/pkg/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserRouter interface {
	HandleUsers(w http.ResponseWriter, req *http.Request)
}

type userRouter struct {
	service service.UserService
}

func NewUserRouter(s service.UserService) *userRouter {
	return &userRouter{s}
}

func (ro *userRouter) HandleUsers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		var err error
		resp := make([]byte, 0)

		emails, found := req.URL.Query()["email"]
		if !found {
			users := ro.service.GetUsers()
			resp, err = json.Marshal(users)
			if err != nil {
				http.Error(w, "error marshaling response", http.StatusInternalServerError)
				return
			}
		} else {
			email := emails[0]
			user := ro.service.GetUser(email)
			resp, err = json.Marshal(user)
			if err != nil {
				http.Error(w, "error marshaling response", http.StatusInternalServerError)
				return
			}
		}

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, "error writing response", http.StatusInternalServerError)
			return
		}
		return
	case http.MethodPost:
		var u *model.User
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			log.Println(fmt.Sprintf("error parsing body content %s", err.Error()))
			http.Error(w, "error parsing body content", http.StatusBadRequest)
			return
		}

		err = ro.service.CreateUser(u)
		if err != nil {
			log.Println(fmt.Sprintf("error creating user %s", err.Error()))
			http.Error(w, "error creating user", http.StatusInternalServerError)
			return
		}

		return
	default:
		http.Error(w, "not implemented", http.StatusNotImplemented)
		return
	}
}
