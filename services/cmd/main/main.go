package main

import (
	authA "github.com/jaax2707/ToDoGorm/services/auth/access"
	authC "github.com/jaax2707/ToDoGorm/services/auth/controllers"
	taskA "github.com/jaax2707/ToDoGorm/services/task/access"
	taskC "github.com/jaax2707/ToDoGorm/services/task/controllers"
	"github.com/jaax2707/ToDoGorm/services/units"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	db := units.InitDB()
	defer db.Close()
	c := cache.New(10*time.Minute, 10*time.Minute)

	mw := units.NewCacheMiddleware(c)

	authA := authA.NewAuthAccess(db)
	taskA := taskA.NewTaskAccess(db)

	task := taskC.NewTask(taskA, c)
	auth := authC.NewAuth(authA, c)

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
