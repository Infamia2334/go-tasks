package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Infamia2334/go-tasks/internal/models"
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

func (taskhandler *TaskHandler) CreateTask(res http.ResponseWriter, req *http.Request) {
	var task models.Task

	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(res, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newTask := taskhandler.taskServices.CreateTask(task)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newTask)
}

func (taskHandler *TaskHandler) GetPaginatedTasks(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	status := query.Get("status")
	assigner := query.Get("assigner")
	assigned_to := query.Get("assigned_to")

	tasks := taskHandler.taskServices.GetPaginatedTasks(page, limit, status, assigner, assigned_to)

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(tasks); err != nil {
		http.Error(res, "Failed to encode response", http.StatusInternalServerError)
	}
}
