package main

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/controllers"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/robfig/go-cache"
	"time"
)

func main() {
	db := InitDB()
	a := access.NewDb(db)
	c := cache.New(5*time.Minute, 5*time.Minute)

	task := controllers.NewDbController(a, c)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/login", task.Login)
	e.POST("/register", task.Register)

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
