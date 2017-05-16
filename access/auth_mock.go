package access

import (
	"errors"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jaax2707/ToDoGorm/utils"
)

// AuthAccessMock represents a struct of DB
type AuthAccessMock struct {
	// db map[string]*models.User
}

// NewAuthAccess return AuthAccessMock object
func NewAuthAccessMock() *AuthAccessMock {
	return &AuthAccessMock{}
}

// CreateUser put User struct into DB and return reference
func (access *AuthAccessMock) CreateUser(u *models.User) models.User {
	return *u
}

// CreateToken create token for User authorization
func (access *AuthAccessMock) CreateToken(username string, password string) (token string, err error) {
	if username == "test11" {
		return "", errors.New("")
	}
	return "token: lskadnlkj", nil
}

// UserExist check if User is in DB table
func (access *AuthAccessMock) UserExist(username string) (user *models.User, err error) {
	if username == "test11" {
		return &models.User{}, errors.New("")
	}
	return &models.User{
		Username: "test",
		Password: utils.Hash([]byte("1111")),
	}, nil
}
