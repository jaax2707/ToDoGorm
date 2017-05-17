package controllers

import (
	"github.com/elithrar/simple-scrypt"
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jaax2707/ToDoGorm/utils"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
)

// Auth represents struct of cache and AuthAccessMock
type Auth struct {
	cache  *cache.Cache
	access access.IAuthAccess
}

// NewAuth return Auth Object
func NewAuth(access access.IAuthAccess, cache *cache.Cache) *Auth {
	return &Auth{access: access, cache: cache}
}

// Login get data from JSON (username, password),
// compare login password and DB password,
// return ErrUnauthorized if it is not equal,
// return StatusOK and token if err == nil
func (ctrl *Auth) Login(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	us, err := ctrl.access.GetUser(u.Username)

	if err != nil || scrypt.CompareHashAndPassword([]byte(us.Password), []byte(u.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, "wrong username or password")
	}
	t, e := ctrl.access.CreateToken(us.Username, us.Password)
	if e != nil {
		return c.JSON(http.StatusUnauthorized, "wrong username or password")
	}
	ctrl.cache.Add(t, "token", cache.DefaultExpiration)
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// Register get data from JSON (username, password),
// if User exist return StatusBadRequest,
// if is not exist put User struct into DB and return StatusOK
func (ctrl *Auth) Register(c echo.Context) error {
	u := models.User{}
	c.Bind(&u)
	u.Password = utils.Hash([]byte(u.Password))
	_, err := ctrl.access.CreateUser(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "this username already exist")
	}
	return c.JSON(http.StatusOK, "register successful")
}
