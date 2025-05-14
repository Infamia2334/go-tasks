package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Infamia2334/go-tasks/internal/models"
	"github.com/Infamia2334/go-tasks/internal/services"
)

type TaskHandler struct {
	taskModel    models.TaskModel
	taskServices services.TaskService
}

func NewTaskHandler(taskModel models.TaskModel, taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskModel:    taskModel,
		taskServices: taskService,
	}
}

func (h *TaskHandler) GetTasks(res http.ResponseWriter, req *http.Request) {
	h.taskModel = models.TaskModel{
		ID:          "1",
		Name:        "Task Uno",
		Description: "This is the first task",
		Status:      "Pending",
		Assignee:    "Alexender Dumas",
		AssignedTo:  "Javert",
		CreatedAt:   "2025-05-14",
		UpdatedAt:   "2025-05-14",
	}
	h.taskServices = services.TaskService{}

	json.NewEncoder(res).Encode(h.taskModel)
}
