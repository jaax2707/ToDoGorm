package models

import (
	"github.com/jinzhu/gorm"
)

// Token define the property - token
type Token struct {
	gorm.Model
	Token string
}

// User defines the properties of User information
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}