package access

import (
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
)

type TaskAccess struct {
	DB *gorm.DB
}

func NewTaskAccess(DB *gorm.DB) *TaskAccess {
	return &TaskAccess{DB}
}


func (access *TaskAccess) GetTask() []models.Task {
	tasks := []models.Task{}
	access.DB.Find(&tasks)
	return tasks
}

func (access *TaskAccess) PostTask(t *models.Task) models.Task {
	defer access.DB.Create(t)
	return *t
}

func (access *TaskAccess) DeleteTask(id string) {
	access.DB.Where("id = ?", id).Delete(&models.Task{})
}
