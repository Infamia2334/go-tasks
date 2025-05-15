package services

import (
	"strconv"
	"time"

	"github.com/Infamia2334/go-tasks/internal/models"
)

type TaskService interface {
	GetTasks() []models.Task
	CreateTask(task models.Task) models.Task
	GetPaginatedTasks(page, limit int, status string, assigner string, assigned_to string) []models.Task
}

type taskServiceStruct struct {
	tasks []models.Task
}

func NewTaskService() TaskService {
	return &taskServiceStruct{
		tasks: []models.Task{
			{
				ID:          "1",
				Name:        "Task Uno",
				Description: "This is the first task",
				Status:      "Pending",
				Assigner:    "Alexender Dumas",
				AssignedTo:  "Javert",
				CreatedAt:   "2025-05-14",
				UpdatedAt:   "2025-05-14",
			},
			{
				ID:          "2",
				Name:        "Task Dos",
				Description: "This is the second task",
				Status:      "In Progress",
				Assigner:    "Victor Hugo",
				AssignedTo:  "Jean Valjean",
				CreatedAt:   "2025-05-15",
				UpdatedAt:   "2025-05-15",
			},
		},
	}
}

func (taskService *taskServiceStruct) GetTasks() []models.Task {
	return taskService.tasks
}

func (taskService *taskServiceStruct) GetPaginatedTasks(page, limit int, status string, assigner string, assigned_to string) []models.Task {
	var filtered []models.Task

	for i, task := range taskService.tasks {
		if i >= (page-1)*limit && i < page*limit {
			if (status == "" || task.Status == status) &&
				(assigner == "" || task.Assigner == assigner) &&
				(assigned_to == "" || task.AssignedTo == assigned_to) {
				filtered = append(filtered, task)
			}
		}
	}

	if len(filtered) == 0 {
		return []models.Task{}
	}

	offset := (page - 1) * limit
	if offset >= len(filtered) {
		return []models.Task{}
	}

	end := offset + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	filtered = filtered[offset:end]
	return filtered
}

func (taskService *taskServiceStruct) CreateTask(task models.Task) models.Task {
	task.ID = strconv.Itoa(len(taskService.tasks) + 1)
	task.CreatedAt = time.Now().Format("2006-01-02")
	task.UpdatedAt = time.Now().Format("2006-01-02")
	taskService.tasks = append(taskService.tasks, task)

	return task
}
