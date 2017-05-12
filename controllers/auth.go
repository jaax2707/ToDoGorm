package controllers

import (
	"github.com/elithrar/simple-scrypt"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"github.com/pmylund/go-cache"
)

func (ctrl *DbController) Login(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	pass := u.Password
	us := ctrl.access.DB.Where("username = ?", u.Username).Find(&u)
	if us.RecordNotFound() == false {
		key := u.Password
		err := scrypt.CompareHashAndPassword([]byte(key), []byte(pass))
		if err == nil {
			t := ctrl.access.Login(u.Username, key)
			ctrl.cache.Add(t, "token", cache.DefaultExpiration)
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		}
	}
	return echo.ErrUnauthorized
}

func (ctrl *DbController) Register(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	us := ctrl.access.DB.Where("username = ?", u.Username).Find(&u)
	if us.RecordNotFound() == false {
		return c.JSON(http.StatusMethodNotAllowed, "this username already exist")
	}
	u.Password = Hash([]byte(u.Password))
	ctrl.access.Register(&u)
	return c.JSON(http.StatusOK, "register successful")
}

func Hash(password []byte) string {
	hash, err := scrypt.GenerateFromPassword(password, scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
