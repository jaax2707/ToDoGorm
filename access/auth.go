package access

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jaax2707/ToDoGorm/models"
	"time"
	"github.com/jinzhu/gorm"
)

type AuthAccess struct {
	DB *gorm.DB
}

func NewAuthAccess(DB *gorm.DB) *AuthAccess {
	return &AuthAccess{DB}
}

func (access *AuthAccess) Register(u *models.User) models.User {
	defer access.DB.Create(&u)
	return *u
}

func (access *AuthAccess) Login(username string, password string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return t
}
