package access

import (
	"errors"
	"github.com/jaax2707/ToDoGorm/models"
)

// AuthAccessMock represents a struct of DB
type AuthAccessMock struct {
	db map[string]*models.User
}

// NewAuthAccess return AuthAccessMock object
func NewAuthAccessMock() *AuthAccessMock {
	return &AuthAccessMock{
		make(map[string]*models.User),
	}
}

// CreateUser put User struct into DB and return reference
func (access *AuthAccessMock) CreateUser(u *models.User) (*models.User, error) {

	x := access.db[u.Username]
	if x != nil {
		return &models.User{}, errors.New("")
	}

	access.db[u.Username] = u
	return u, nil
}

// CreateToken create token for User authorization
func (access *AuthAccessMock) CreateToken(username string, password string) (token string, err error) {
	return "", nil
}

// GetUser check if User is in DB table
func (access *AuthAccessMock) GetUser(username string) (user *models.User, err error) {
	user = &models.User{}
	x := access.db[username]
	if x != nil {
		return x, nil
	}
	return &models.User{}, errors.New("")
}
