package controllers

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
)

// Task represents struct of cache and AuthAccessMock
type Task struct {
	cache  *cache.Cache
	access *access.TaskAccess
}

// NewTask return Task Object
func NewTask(access *access.TaskAccess, cache *cache.Cache) *Task {
	return &Task{access: access, cache: cache}
}

// GetAll return GetTasks method and StatusOK
func (ctrl *Task) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.access.GetTasks())
}

// PostTask get data from JSON (name),
// call PutTask method and return StatusCreated
func (ctrl *Task) PostTask(c echo.Context) error {
	task := models.Task{}
	c.Bind(&task)
	ctrl.access.PutTask(&task)
	name := task.Name
	return c.JSON(http.StatusCreated, "created: "+name)
}

// DeleteTask get data from param (id),
// call DeleteTask method and return StatusOK
func (ctrl *Task) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	ctrl.access.DeleteTask(id)
	return c.String(http.StatusOK, "Deleted: "+id)
}
