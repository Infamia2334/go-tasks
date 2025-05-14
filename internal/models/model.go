package models

type TaskModel struct {
	ID          string
	Name        string
	Description string
	Status      string
	Assignee    string
	AssignedTo  string
	CreatedAt   string
	UpdatedAt   string
}
