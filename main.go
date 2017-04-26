package main
import(
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)
func main(){

	e := echo.New()
	e.POST("/task", SaveTask)
	e.GET("/task", GetTask)
	e.PATCH("/task/:id", DeleteTask)
	e.Start(":8000")
}

type Task struct {
	gorm.Model
	Name string `json:"name"`
}

func SaveTask(c echo.Context) error {
	db := InitDB()
	if db.HasTable(&Task{}) == false {
		db.CreateTable(&Task{})
	}
	t := Task{}
	c.Bind(&t)
	name := t.Name
	db.Create(&t)
	defer db.Close()
	return c.JSON(http.StatusCreated, "created: " + name)
}

func GetTask(c echo.Context) error {
	db := InitDB()
	tasks := []Task{}
	allTasks := db.Find(&tasks)
	defer db.Close()
	return c.JSON(http.StatusOK, allTasks)
}
func DeleteTask (c echo.Context) error {
	db := InitDB()
	id := c.Param("id")
	db.Where("id = ?", id).Delete(&Task{})
	defer db.Close()
	return c.String(http.StatusOK, "Deleted: " + id)
}

func InitDB() (*gorm.DB) {
	db, err := gorm.Open("postgres", "user=postgres password=2707 dbname=owner sslmode=disable")
	if err != nil {
		panic("failed to connect database !!")
	}
	return db
}