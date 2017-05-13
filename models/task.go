package models

import (
	"github.com/jinzhu/gorm"
)

// Task represents struct of Task Name
type Task struct {
	gorm.Model
	Name string `json:"name"`
}
