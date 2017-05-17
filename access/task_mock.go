package access

import "github.com/jaax2707/ToDoGorm/models"

// TaskAccess represents a struct of DB
type TaskAccessMock struct {
}

// NewTaskAccess return TaskAccess object
func NewTaskAccessMock() *TaskAccess {
	return &TaskAccess{}
}

// GetTasks find tasks into DB and return Task objects
func (access *TaskAccessMock) GetTasks() []models.Task {
	return []models.Task{}
}

// PutTask put Task struct into DN and return reference
func (access *TaskAccessMock) PutTask(t *models.Task) models.Task {
	return *t
}

// DeleteTask fund Task ID into DB and delete Task
func (access *TaskAccessMock) DeleteTask(id string) {
}
