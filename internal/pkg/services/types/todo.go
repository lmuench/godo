package types

import "github.com/jinzhu/gorm"

// Todo type
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
