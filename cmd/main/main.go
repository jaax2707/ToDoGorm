package main

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/controllers"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	db := InitDB()
	defer db.Close()
	c := cache.New(5*time.Minute, 5*time.Minute)

	authA := access.NewAuthAccess(db)
	taskA := access.NewTaskAccess(db)

	task := controllers.NewTaskController(taskA, c)
	auth := controllers.NewAuthController(authA, c)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/login", auth.Login)
	e.POST("/register", auth.Register)

	r := e.Group("/restricted")
	r.POST("/task", task.PostTask)
	r.GET("/task", task.GetAll)
	r.PATCH("/task/:id", task.DeleteTask)

	e.Logger.Fatal(e.Start(":8000"))
}

func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=2707 dbname=owner sslmode=disable")
	db.AutoMigrate(models.Task{}, models.User{})
	if err != nil {
		panic("failed to connect database !!")
	}
	return db
}
