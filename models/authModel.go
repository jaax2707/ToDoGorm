package models

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

type Token struct {
	Token string
}

func (access *TaskDataAccess) Register (u *User) User {
	defer access.DB.Create(&u)
	return *u
}

func (access *TaskDataAccess) Login (username string, password string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Printf("%T", token)
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