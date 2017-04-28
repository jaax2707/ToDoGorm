package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/elithrar/simple-scrypt"
	"log"
)

type TaskController struct {
	access *models.TaskDataAccess
}

func NewTaskController (access *models.TaskDataAccess) *TaskController {
	return &TaskController{access}
}

func (ctrl *TaskController) GetAll (c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.access.GetTask())
}
func (ctrl *TaskController) PostTask (c echo.Context) error {
	task := models.Task{}
	c.Bind(&task)
	ctrl.access.PostTask(&task)
	name := task.Name
	return c.JSON(http.StatusCreated, "created: " + name)
}
func (ctrl *TaskController) DeleteTask (c echo.Context) error {
	id  := c.Param("id")
	ctrl.access.DeleteTask(id)
	return c.String(http.StatusOK, "Deleted: " + id)
}
func (ctrl *TaskController) Login (c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	pass := u.Password
	us := ctrl.access.DB.Where("username = ?", u.Username).Find(&u)
	if us.RecordNotFound() == false {
		key := u.Password
		err := scrypt.CompareHashAndPassword([]byte(key), []byte(pass))
		if err == nil{
			ctrl.access.Login(u.Username, key)
			return c.JSON(http.StatusOK, "login succesful")
		}
	}
	return echo.ErrUnauthorized
}
func (ctrl *TaskController) Register (c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	u.Password = Hash([]byte(u.Password))
	ctrl.access.Register(&u)
	return c.JSON(http.StatusOK, "register successful")
}
func Hash (password []byte) string {
	hash, err := scrypt.GenerateFromPassword(password, scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
