package models

import "github.com/jinzhu/gorm"

type TaskDataAccess struct {
	DB *gorm.DB
}

func NewTaskDataAccess (DB *gorm.DB) *TaskDataAccess {
	return &TaskDataAccess{DB}
}
