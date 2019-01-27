package models

import (
	"github.com/jinzhu/gorm"
)

// Todo model
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetAll todos
func (repo *Repo) GetAll() []Todo {
	var todos []Todo
	repo.DB.Find(&todos)
	return todos
}

// Repo - Todo Repository
type Repo struct {
	DB *gorm.DB
}
