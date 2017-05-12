package models

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	Token string
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
