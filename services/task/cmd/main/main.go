package main

import (
	pb "github.com/jaax2707/ToDoGorm/services/auth/protobuf"
	taskA "github.com/jaax2707/ToDoGorm/services/task/access"
	taskC "github.com/jaax2707/ToDoGorm/services/task/controllers"
	"github.com/jaax2707/ToDoGorm/services/task/units"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	db := units.InitDB()
	defer db.Close()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTokenClient(conn)
	mw := units.NewClient(client)

	taskA := taskA.NewTaskAccess(db)
	task := taskC.NewTask(taskA)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/restricted")
	r.Use(mw.CheckToken)
	r.POST("/task", task.PostTask)
	r.GET("/task", task.GetAll)
	r.PATCH("/task/:id", task.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
