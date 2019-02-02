package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"AUTO_INCREMENT"`
	Password string `json:"password"`
}

// CreateUser creates new user record
func (repo UserRepo) CreateUser(user User) error {
	return repo.DB.Create(&user).Error
}

// UsernameTaken returns true if the provided username is already taken
func (repo UserRepo) UsernameTaken(username string) bool {
	return !repo.DB.Where("username = ?", username).First(&User{}).RecordNotFound()
}

// UserRepo repository
type UserRepo struct {
	DB *gorm.DB
}
