package controllers

import (
	"github.com/elithrar/simple-scrypt"
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
)

// Auth represents struct of cache and AuthAccess
type Auth struct {
	cache  *cache.Cache
	access *access.AuthAccess
}

// NewAuth return Auth Object
func NewAuth(access *access.AuthAccess, cache *cache.Cache) *Auth {
	return &Auth{access: access, cache: cache}
}

// Login get data from JSON (username, password),
// compare login password and DB password,
// return ErrUnauthorized if it is not equal,
// return StatusOK and token if err == nil
func (ctrl *Auth) Login(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	pass := u.Password
	if ctrl.access.UserExist(&u) {
		key := u.Password
		err := scrypt.CompareHashAndPassword([]byte(key), []byte(pass))
		if err == nil {
			t := ctrl.access.CreateToken(u.Username, key)
			ctrl.cache.Add(t, "token", cache.DefaultExpiration)
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		}
	}
	return echo.ErrUnauthorized
}

// Register get data from JSON (username, password),
// if User exist return StatusMethodNotAllowed,
// if is not exist put User struct into DB and return StatusOK
func (ctrl *Auth) Register(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	if ctrl.access.UserExist(&u) {
		return c.JSON(http.StatusMethodNotAllowed, "this username already exist")
	}
	u.Password = Hash([]byte(u.Password))
	ctrl.access.CreateUser(&u)
	return c.JSON(http.StatusOK, "register successful")
}

// Hash create and return hash from given password
func Hash(password []byte) string {
	hash, err := scrypt.GenerateFromPassword(password, scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
