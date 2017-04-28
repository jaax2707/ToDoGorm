package main
import(
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/jaax2707/ToDoGorm/controllers"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main(){
	db := InitDB()
	access := models.NewTaskDataAccess(db)
	task := controllers.NewTaskController(access)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/login", task.Login)
	e.POST("/register", task.Register)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("/task", task.PostTask)
	r.GET("/task", task.GetAll)
	r.PATCH("/task/:id", task.DeleteTask)

	e.Logger.Fatal(e.Start(":8000"))
}

func InitDB() (*gorm.DB) {

	db, err := gorm.Open("postgres", "user=postgres password=2707 dbname=owner sslmode=disable")
	if db.HasTable(models.Task{}) == false {
		db.CreateTable(models.Task{})
	}
	if db.HasTable(models.User{}) == false {
		db.CreateTable(models.User{})
	}
	if err != nil {
		panic("failed to connect database !!")
	}
	return db
}
