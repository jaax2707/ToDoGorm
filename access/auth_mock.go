package access

import (
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jaax2707/ToDoGorm/utils"
)

// AuthAccessMock represents a struct of DB
type AuthAccessMock struct {
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
	return "", nil
}

// UserExist check if User is in DB table
func (access *AuthAccessMock) GetUser(username string) (user *models.User, err error) {
	return &models.User{
		Username: "test",
		Password: utils.Hash([]byte("1111")),
	}, nil
}
