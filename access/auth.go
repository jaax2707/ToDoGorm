package access

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jinzhu/gorm"
	"time"
)

// AuthAccess represents a struct of DB
type AuthAccess struct {
	DB *gorm.DB
}

// NewAuthAccess return AuthAccess object
func NewAuthAccess(DB *gorm.DB) *AuthAccess {
	return &AuthAccess{DB}
}

// CreateUser put User struct into DB and return reference
func (access *AuthAccess) CreateUser(u *models.User) models.User {
	defer access.DB.Create(&u)
	return *u
}

// CreateToken create token for User authorization
func (access *AuthAccess) CreateToken(username string, password string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return t
}

// UserExist check if User is in DB table
func (access *AuthAccess) UserExist(user *models.User) bool {
	us := access.DB.Where("username = ?", user.Username).Find(&user)
	if !us.RecordNotFound() {
		return true
	}
	return false
}
