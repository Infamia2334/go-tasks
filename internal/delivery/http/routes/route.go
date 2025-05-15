package route

import (
	handlers "github.com/Infamia2334/go-tasks/internal/delivery/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, taskHandler *handlers.TaskHandler) {
	router.HandleFunc("/tasks/all", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/task", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", taskHandler.GetPaginatedTasks).Methods("GET")
}
