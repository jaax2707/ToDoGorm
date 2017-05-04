package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/elithrar/simple-scrypt"
	"log"
)

func (ctrl *TaskController) Login (c echo.Context) error {
	u := models.User{}
	t := models.Token{}
	c.Bind(&u)
	pass := u.Password
	us := ctrl.access.DB.Where("username = ?", u.Username).Find(&u)
	if us.RecordNotFound() == false {
		key := u.Password
		err := scrypt.CompareHashAndPassword([]byte(key), []byte(pass))
		if err == nil{
			t.Token = ctrl.access.Login(u.Username, key)
			return c.JSON(http.StatusOK, "login succesful")
		}
	}
	return echo.ErrUnauthorized
}

func (ctrl *TaskController) Register (c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	us := ctrl.access.DB.Where("username = ?", u.Username).Find(&u)
	if us.RecordNotFound() == false{
		return c.JSON(http.StatusMethodNotAllowed, "this username already exist")
	}
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
