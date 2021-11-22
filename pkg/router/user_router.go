package router

import (
	"elon-task/pkg/service"
	"encoding/json"
	"log"
	"net/http"
)

type UserRouter interface {
	HandleUsers(w http.ResponseWriter, req *http.Request)
}

type userRouter struct {
	service service.UserService
	logger  log.Logger
}

func NewRouter(s service.UserService, l log.Logger) *userRouter {
	return &userRouter{s, l}
}

func (ro *userRouter) HandleUsers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		users := ro.service.GetUsers()

		resp, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "error marshaling repsonse", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, "error writing response", http.StatusInternalServerError)
			return
		}
	}
}
