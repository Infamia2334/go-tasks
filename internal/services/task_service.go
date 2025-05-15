package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Infamia2334/go-tasks/internal/models"
)

type TaskService interface {
	GetTasks() []models.Task
	CreateTask(task models.Task) models.Task
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

func (taskService *taskServiceStruct) CreateTask(task models.Task) models.Task {
	fmt.Printf("Creating task: %+v\n", task)
	task.ID = strconv.Itoa(len(taskService.tasks) + 1)
	task.CreatedAt = time.Now().Format("2006-01-02")
	task.UpdatedAt = time.Now().Format("2006-01-02")
	taskService.tasks = append(taskService.tasks, task)

	return task
}
