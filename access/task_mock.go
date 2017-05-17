package access

import (
	"errors"
	"github.com/jaax2707/ToDoGorm/models"
	"strconv"
)

// TaskAccess represents a struct of DB
type TaskAccessMock struct {
	db map[int]*models.Task
}

// NewTaskAccess return TaskAccess object
func NewTaskAccessMock() *TaskAccessMock {
	return &TaskAccessMock{make(map[int]*models.Task)}
}

// GetTasks find tasks into DB and return Task objects
func (access *TaskAccessMock) GetTasks() (tasks *[]models.Task, err error) {
	return
}

// PutTask put Task struct into DN and return reference
func (access *TaskAccessMock) PutTask(t *models.Task) error {
	if t.Name != "" {
		access.db[len(access.db)+1] = t
		return nil
	}
	return errors.New("")
}

// DeleteTask find Task ID into DB and delete Task
func (access *TaskAccessMock) DeleteTask(id string) error {
	x, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	it := access.db[x]

	if it != nil {
		delete(access.db, x)
		return nil

	}
	return errors.New("")
}
