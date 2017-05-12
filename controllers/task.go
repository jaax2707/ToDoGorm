package controllers

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
)

type TaskController struct {
	cache  *cache.Cache
	access *access.TaskAccess
}

func NewTaskController(access *access.TaskAccess, cache *cache.Cache) *TaskController {
	return &TaskController{access: access, cache: cache}
}

func (ctrl *TaskController) GetAll(c echo.Context) error {
	t := c.Request().Header.Get("Authorization")
	_, exist := ctrl.cache.Get(t)
	if !exist {
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, ctrl.access.GetTask())
}

func (ctrl *TaskController) PostTask(c echo.Context) error {
	t := c.Request().Header.Get("Authorization")
	_, exist := ctrl.cache.Get(t)
	if !exist {
		return c.NoContent(http.StatusUnauthorized)
	}
	task := models.Task{}
	c.Bind(&task)
	ctrl.access.PostTask(&task)
	name := task.Name
	return c.JSON(http.StatusCreated, "created: "+name)
}

func (ctrl *TaskController) DeleteTask(c echo.Context) error {
	t := c.Request().Header.Get("Authorization")
	_, exist := ctrl.cache.Get(t)
	if !exist {
		return c.NoContent(http.StatusUnauthorized)
	}
	id := c.Param("id")
	ctrl.access.DeleteTask(id)
	return c.String(http.StatusOK, "Deleted: "+id)
}
