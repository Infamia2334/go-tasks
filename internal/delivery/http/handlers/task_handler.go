package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Infamia2334/go-tasks/internal/services"
)

type TaskHandler struct {
	taskServices services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskServices: taskService,
	}
}

func (taskHandler *TaskHandler) GetTasks(res http.ResponseWriter, req *http.Request) {
	tasks := taskHandler.taskServices.GetTasks()

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(tasks); err != nil {
		http.Error(res, "Failed to encode response", http.StatusInternalServerError)
	}
}
