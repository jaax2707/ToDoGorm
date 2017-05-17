package access

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
	"time"
)

// AuthAccessMock represents a struct of DB
type AuthAccess struct {
	DB *gorm.DB
}

type IAuthAccess interface {
	CreateUser(u *models.User) models.User
	CreateToken(username string, password string) (token string, err error)
	GetUser(username string) (user *models.User, err error)
}

// NewAuthAccess return AuthAccessMock object
func NewAuthAccess(DB *gorm.DB) *AuthAccess {
	return &AuthAccess{DB}
}

// CreateUser put User struct into DB and return reference
func (access *AuthAccess) CreateUser(u *models.User) models.User {
	defer access.DB.Create(&u)
	return *u
}

// CreateToken create token for User authorization
func (access *AuthAccess) CreateToken(username string, password string) (token string, err error) {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token, err = t.SignedString([]byte("secret"))
	return
}

// GetUser check if User is in DB table
func (access *AuthAccess) GetUser(username string) (user *models.User, err error) {
	user = &models.User{}
	err = access.DB.Where("username = ?", username).Find(user).Error
	return
}
