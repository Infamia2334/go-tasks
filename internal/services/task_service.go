package services

import "github.com/Infamia2334/go-tasks/internal/models"

type TaskService interface {
	GetTasks() []models.Task
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
				Assignee:    "Alexender Dumas",
				AssignedTo:  "Javert",
				CreatedAt:   "2025-05-14",
				UpdatedAt:   "2025-05-14",
			},
			{
				ID:          "2",
				Name:        "Task Dos",
				Description: "This is the second task",
				Status:      "In Progress",
				Assignee:    "Victor Hugo",
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
