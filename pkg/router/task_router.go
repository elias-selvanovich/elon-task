package router

import (
	"elon-task/pkg/model"
	"elon-task/pkg/service"
	"encoding/json"
	"net/http"
)

type TaskRouter interface {
	HandleTasks(w http.ResponseWriter, req *http.Request)
}

type taskRouter struct {
	taskService service.TaskService
}

func NewTaskRouter(s service.TaskService) *taskRouter {
	return &taskRouter{s}
}

func (r *taskRouter) HandleTasks(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		emails, found := req.URL.Query()["email"]
		if !found || len(emails) < 1 {
			http.Error(w, "email is required", http.StatusBadRequest)
			return
		}

		tasks, err := r.taskService.GetTasks(emails[0])
		if err != nil {
			http.Error(w, "error getting tasks", http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "error marshalling response", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, "error writing response", http.StatusInternalServerError)
			return
		}

		return
	case http.MethodPost:
		var t *model.Task
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			http.Error(w, "error parsing body", http.StatusBadRequest)
			return
		}

		err = r.taskService.CreateTask(t)
		if err != nil {
			http.Error(w, "error creating task", http.StatusInternalServerError)
			return
		}
		return
	case http.MethodPut:
		var t *model.Task
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			http.Error(w, "error parsing body", http.StatusBadRequest)
			return
		}

		err = r.taskService.UpdateTask(t)
		if err != nil {
			http.Error(w, "error updating task", http.StatusInternalServerError)
			return
		}
		return
	case http.MethodDelete:
		var t *model.Task
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			http.Error(w, "error parsing body", http.StatusBadRequest)
			return
		}

		err = r.taskService.DeleteTask(t)
		if err != nil {
			http.Error(w, "error updating task", http.StatusInternalServerError)
			return
		}
		return
	default:
		http.Error(w, "not implemented", http.StatusNotImplemented)
		return
	}
}
