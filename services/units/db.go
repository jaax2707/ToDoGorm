package units

import (
	auth "github.com/jaax2707/ToDoGorm/services/auth/models"
	task "github.com/jaax2707/ToDoGorm/services/task/models"
	"github.com/jinzhu/gorm"
)

// InitDB Open postgres DB
func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=2707 dbname=owner sslmode=disable")
	db.AutoMigrate(task.Task{}, auth.User{})
	if err != nil {
		panic("failed to connect database !!")
	}
	return db
}
