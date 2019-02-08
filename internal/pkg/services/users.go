package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lmuench/godo/internal/pkg/services/types"
	"golang.org/x/crypto/bcrypt"
)

// createUser creates new user record
func (s Users) createUser(user types.User) error {
	return s.DB.Create(&user).Error
}

// SignUp validates username and password lengths, hashes password and calls CreateUser
func (s Users) SignUp(_user types.User) error {
	if len(_user.Username) < 3 {
		return errors.New("Username must be at least 3 characters long")
	}

	if len(_user.Password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	if s.UsernameTaken(_user.Username) {
		return errors.New("Username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(_user.Password), 8)
	if err != nil {
		return errors.New("Please choose a different password")
	}

	user := types.User{
		Username: _user.Username,
		Password: string(hashedPassword),
	}
	err = s.createUser(user)
	if err != nil {
		return errors.New("Please choose a different username")
	}
	return nil
}

// GetUser returns user with provided username
func (s Users) GetUser(username string) (types.User, error) {
	var user types.User
	if s.DB.Where("username = ?", username).First(&user).RecordNotFound() {
		return user, errors.New("User not found")
	}
	return user, nil
}

// UsernameTaken returns true if the provided username is already taken
func (s Users) UsernameTaken(username string) bool {
	return !s.DB.Where("username = ?", username).First(&types.User{}).RecordNotFound()
}

// Users ...
type Users struct {
	DB *gorm.DB
}
