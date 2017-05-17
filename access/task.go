package access

import (
"errors"
"github.com/jaax2707/ToDoGorm/models"
"github.com/jinzhu/gorm"
)

// TaskAccess represents a struct of DB
type TaskAccess struct {
	DB *gorm.DB
}

type ITaskAccess interface {
	GetTasks() (tasks *[]models.Task, err error)
	PutTask(t *models.Task) error
	DeleteTask(id string) error
}

// NewTaskAccess return TaskAccess object
func NewTaskAccess(DB *gorm.DB) *TaskAccess {
	return &TaskAccess{DB}
}

// GetTasks find tasks into DB and return Task objects
func (access *TaskAccess) GetTasks() (tasks *[]models.Task, err error) {
	tasks = &[]models.Task{}
	err = access.DB.Find(&tasks).Error
	return
}

// PutTask put Task struct into DN and return reference
func (access *TaskAccess) PutTask(t *models.Task) error {
	if t.Name != "" {
		access.DB.Create(t)
		return nil
	}
	return errors.New("Task name is empty")
}

// DeleteTask fund Task ID into DB and delete Task
func (access *TaskAccess) DeleteTask(id string) error {
	return access.DB.Where("id = ?", id).First(&models.Task{}).Delete(&models.Task{}).Error
}
