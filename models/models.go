package models

import (
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	"time"
)
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	gorm.Model
	Name string `json:"name"`
}
type TasksCollection struct {
	Tasks []Task
}
type TaskDataAccess struct {
	DB *gorm.DB
}
func NewTaskDataAccess (DB *gorm.DB) *TaskDataAccess {
	return &TaskDataAccess{DB}
}
func (access *TaskDataAccess)GetTask () []Task {
	tasks := []Task{}
	access.DB.Find(&tasks)
	return tasks
}
func (access *TaskDataAccess) PostTask (t *Task) Task {
	defer access.DB.Create(t)
	return *t
}
func (access *TaskDataAccess) DeleteTask (id string){
	access.DB.Where("id = ?", id).Delete(&Task{})
}
func (access *TaskDataAccess) Register (u *User) User{
	defer access.DB.Create(&u)
	return *u
}
func (access *TaskDataAccess) Login (username string, password string) error {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	println(t)
	return err
}