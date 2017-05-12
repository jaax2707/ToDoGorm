package main

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/jaax2707/ToDoGorm/controllers"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jaax2707/ToDoGorm/units"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/labstack/echo/middleware"
	"time"
	"github.com/patrickmn/go-cache"
)

func main() {
	db := InitDB()
	defer db.Close()
	c := cache.New(10*time.Minute, 10*time.Minute)

	mw := units.NewCacheMiddleware(c)

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
	r.Use(mw.CheckToken)
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
