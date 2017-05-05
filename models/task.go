package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name string `json:"name"`
}

type TasksCollection struct {
	Tasks []Task
}

