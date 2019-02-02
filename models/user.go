package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
}

// createUser creates new user record
func (repo UserRepo) createUser(user User) error {
	return repo.DB.Create(&user).Error
}

// SignUp validates username and password lengths, hashes password and calls CreateUser
func (repo UserRepo) SignUp(_user User) error {
	if len(_user.Username) < 3 {
		return errors.New("Username must be at least 3 characters long")
	}

	if len(_user.Password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	if repo.UsernameTaken(_user.Username) {
		return errors.New("Username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(_user.Password), 8)
	if err != nil {
		return errors.New("Please choose a different password")
	}

	user := User{
		Username: _user.Username,
		Password: string(hashedPassword),
	}
	err = repo.createUser(user)
	if err != nil {
		return errors.New("Please choose a different username")
	}
	return nil
}

// GetUser returns user with provided username
func (repo UserRepo) GetUser(username string) (User, error) {
	var user User
	if repo.DB.Where("username = ?", username).First(&user).RecordNotFound() {
		return user, errors.New("User not found")
	}
	return user, nil
}

// UsernameTaken returns true if the provided username is already taken
func (repo UserRepo) UsernameTaken(username string) bool {
	return !repo.DB.Where("username = ?", username).First(&User{}).RecordNotFound()
}

// UserRepo repository
type UserRepo struct {
	DB *gorm.DB
}
