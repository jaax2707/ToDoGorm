package access

import (
	"errors"
	"github.com/jaax2707/ToDoGorm/models"
)

// TaskAccess represents a struct of DB
type TaskAccessMock struct {
}

// NewTaskAccess return TaskAccess object
func NewTaskAccessMock() *TaskAccessMock {
	return &TaskAccessMock{}
}

// GetTasks find tasks into DB and return Task objects
func (access *TaskAccessMock) GetTasks() (tasks *[]models.Task, err error) {
	return
}

// PutTask put Task struct into DN and return reference
func (access *TaskAccessMock) PutTask(t *models.Task) error {
	if t.Name != "" {
		return nil
	}
	return errors.New("")
}

// DeleteTask fund Task ID into DB and delete Task
func (access *TaskAccessMock) DeleteTask(id string) error {
	if id == "3" {
		return errors.New("")
	}
	return nil
}
