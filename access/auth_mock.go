package access

import (
	"errors"
	"github.com/jaax2707/ToDoGorm/models"
	"github.com/jaax2707/ToDoGorm/utils"
)

// AuthAccessMock represents a struct of DB
type AuthAccessMock struct {
	db map[string]*models.User
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
func (access *AuthAccessMock) UserExist(username string) (user *models.User, err error) {
	access.db = make(map[string]*models.User)

	user = &models.User{
		Username: "test",
		Password: utils.Hash([]byte("1111")),
	}

	access.db[user.Username] = user
	chkUser := access.db[username]
	if chkUser != nil {
		return user, nil
	}
	return &models.User{}, errors.New("")
}
