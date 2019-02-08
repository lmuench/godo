package types

import "github.com/jinzhu/gorm"

// User type
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
}
