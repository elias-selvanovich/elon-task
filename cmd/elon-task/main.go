package main

import (
	"elon-task/pkg/router"
	"elon-task/pkg/service"
	"elon-task/pkg/storage"
	"log"
	"net/http"
)

func main() {
	// init storage
	userStorage := storage.NewUserStorage()
	taskStorage := storage.NewTaskStorage()

	// init service
	userService := service.NewUserService(userStorage)
	taskService := service.NewTaskService(taskStorage)

	// init routers
	userRouter := router.NewUserRouter(userService)
	taskRouter := router.NewTaskRouter(taskService)

	// register endpoints
	http.HandleFunc("/user", userRouter.HandleUsers)
	http.HandleFunc("/task", taskRouter.HandleTasks)

	log.Println("Server starting")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
