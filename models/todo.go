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
func (repo *TodoRepo) GetAll() []Todo {
	var todos []Todo
	repo.DB.Find(&todos)
	return todos
}

// TodoRepo - Todo Repository
type TodoRepo struct {
	DB *gorm.DB
}
