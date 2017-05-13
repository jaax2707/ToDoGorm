package access

import (
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
)

// TaskAccess represents a struct of DB
type TaskAccess struct {
	DB *gorm.DB
}

// NewTaskAccess return TaskAccess object
func NewTaskAccess(DB *gorm.DB) *TaskAccess {
	return &TaskAccess{DB}
}

// GetTasks find tasks into DB and return Task objects
func (access *TaskAccess) GetTasks() []models.Task {
	tasks := []models.Task{}
	access.DB.Find(&tasks)
	return tasks
}

// PutTask put Task struct into DN and return reference
func (access *TaskAccess) PutTask(t *models.Task) models.Task {
	defer access.DB.Create(t)
	return *t
}

// DeleteTask fund Task ID into DB and delete Task
func (access *TaskAccess) DeleteTask(id string) {
	access.DB.Where("id = ?", id).Delete(&models.Task{})
}
