package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
}

// CreateUser creates new user record
func (repo UserRepo) CreateUser(user User) error {
	return repo.DB.Create(&user).Error
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
