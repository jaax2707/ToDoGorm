package main
import(
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/jaax2707/ToDoGorm/controllers"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/labstack/echo"
)

func main(){
	access := models.NewTaskDataAccess(InitDB())
	task := controllers.NewTaskController(access)
	defer InitDB().Close()

	e := echo.New()
	e.POST("/task", task.PostTask)
	e.GET("/task", task.GetAll)
	e.PATCH("/task/:id", task.DeleteTask)
	e.Start(":8000")
}

type Task struct {
	gorm.Model
	Name string `json:"name"`
}

func InitDB() (*gorm.DB) {
	db, err := gorm.Open("postgres", "user=postgres password=2707 dbname=owner sslmode=disable")
	if err != nil {
		panic("failed to connect database !!")
	}
	return db
}