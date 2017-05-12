package controllers

import (
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"net/http"
)

func (ctrl *DbController) GetAll(c echo.Context) error {
	t := c.Request().Header.Get("Authorization")
	_, exist := ctrl.cache.Get(t)
	if !exist {
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, ctrl.access.GetTask())
}

func (ctrl *DbController) PostTask(c echo.Context) error {
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

func (ctrl *DbController) DeleteTask(c echo.Context) error {
	t := c.Request().Header.Get("Authorization")
	_, exist := ctrl.cache.Get(t)
	if !exist {
		return c.NoContent(http.StatusUnauthorized)
	}
	id := c.Param("id")
	ctrl.access.DeleteTask(id)
	return c.String(http.StatusOK, "Deleted: "+id)
}
