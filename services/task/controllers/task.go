package controllers

import (
	"github.com/jaax2707/ToDoGorm/services/task/access"
	"github.com/jaax2707/ToDoGorm/services/task/models"
	"github.com/labstack/echo"
	"net/http"
)

// Task represents struct of cache and AuthAccessMock
type Task struct {
	access access.ITaskAccess
}

// NewTask return Task Object
func NewTask(access access.ITaskAccess) *Task {
	return &Task{access: access}
}

// GetAll return GetTasks method and StatusOK
func (ctrl *Task) GetAll(c echo.Context) error {
	t, err := ctrl.access.GetTasks()
	if err != nil {
		return c.JSON(http.StatusNoContent, "No tasks in db")
	}
	return c.JSON(http.StatusOK, t)
}

// PostTask get data from JSON (name),
// call PutTask method and return StatusBadRequest
func (ctrl *Task) PostTask(c echo.Context) error {
	task := models.Task{}
	c.Bind(&task)
	err := ctrl.access.PutTask(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "task name is empty")
	}
	return c.JSON(http.StatusOK, "created: "+task.Name)
}

// DeleteTask get data from param (id),
// call DeleteTask method and return StatusOK
func (ctrl *Task) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	err := ctrl.access.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Task with ID: "+id+" is not exist")
	}
	return c.String(http.StatusOK, "Deleted: "+id)
}
