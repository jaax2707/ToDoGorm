package access

import (
	"github.com/jaax2707/ToDoGorm/models"
)

func (access *Db) GetTask() []models.Task {
	tasks := []models.Task{}
	access.DB.Find(&tasks)
	return tasks
}

func (access *Db) PostTask(t *models.Task) models.Task {
	defer access.DB.Create(t)
	return *t
}

func (access *Db) DeleteTask(id string) {
	access.DB.Where("id = ?", id).Delete(&models.Task{})
}
