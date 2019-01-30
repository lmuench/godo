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

// GetAllTodos returns all todos
func (repo TodoRepo) GetAllTodos() []Todo {
	var todos []Todo
	repo.DB.Find(&todos)
	return todos
}

// TodoRepo repository
type TodoRepo struct {
	DB *gorm.DB
}
