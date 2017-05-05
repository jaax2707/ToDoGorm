package controllers

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"net/http"
)

type TaskController struct {
	access *access.Db
}

func NewTaskController(access *access.Db) *TaskController {
	return &TaskController{access}
}

func (ctrl *TaskController) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.access.GetTask())
}

func (ctrl *TaskController) PostTask(c echo.Context) error {
	task := models.Task{}
	c.Bind(&task)
	ctrl.access.PostTask(&task)
	name := task.Name
	return c.JSON(http.StatusCreated, "created: "+name)
}

func (ctrl *TaskController) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	ctrl.access.DeleteTask(id)
	return c.String(http.StatusOK, "Deleted: "+id)
}
