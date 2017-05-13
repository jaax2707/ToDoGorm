package units

import (
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
)

// InitDB Open postgres DB
func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=2707 dbname=owner sslmode=disable")
	db.AutoMigrate(models.Task{}, models.User{})
	if err != nil {
		panic("failed to connect database !!")
	}
	return db
}
