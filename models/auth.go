package models

import (
	"github.com/jinzhu/gorm"
)

// Token represents struct of Token
type Token struct {
	gorm.Model
	Token string
}

// User represents struct of User Username and Password
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}