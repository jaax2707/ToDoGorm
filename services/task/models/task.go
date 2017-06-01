package models

import (
	"github.com/jinzhu/gorm"
)

// Task define the property of task name
type Task struct {
	gorm.Model
	Name string `json:"name"`
}
