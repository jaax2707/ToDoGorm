package main

import (
	authA "github.com/jaax2707/ToDoGorm/services/auth/access"
	authC "github.com/jaax2707/ToDoGorm/services/auth/controllers"
	pb "github.com/jaax2707/ToDoGorm/services/auth/protobuf"
	u "github.com/jaax2707/ToDoGorm/services/auth/units"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
	"time"
)

const (
	GRPCport = ":50051"
)

func main() {
	db := u.InitDB()
	defer db.Close()

	c := cache.New(10*time.Minute, 10*time.Minute)

	authA := authA.NewAuthAccess(db)
	auth := authC.NewAuth(authA, c)

	server := authC.NewServer(c)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		lis, err := net.Listen("tcp", GRPCport)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterTokenServer(s, server)
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		wg.Done()
	}()

	go func() {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.POST("/login", auth.Login)
		e.POST("/register", auth.Register)

		e.Logger.Fatal(e.Start(":8000"))

		wg.Done()
	}()

	wg.Wait()
}
