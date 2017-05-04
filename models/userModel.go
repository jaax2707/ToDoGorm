package models

import (
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	gorm.Model
	Name string `json:"name"`
}

type TasksCollection struct {
	Tasks []Task
}

type TaskDataAccess struct {
	DB *gorm.DB
}

func NewTaskDataAccess (DB *gorm.DB) *TaskDataAccess {
	return &TaskDataAccess{DB}
}

func (access *TaskDataAccess)GetTask () []Task {
	tasks := []Task{}
	access.DB.Find(&tasks)
	return tasks
}

func (access *TaskDataAccess) PostTask (t *Task) Task {
	defer access.DB.Create(t)
	return *t
}

func (access *TaskDataAccess) DeleteTask (id string) {
	access.DB.Where("id = ?", id).Delete(&Task{})
}