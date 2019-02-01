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

// GetTodos returns all todos
func (repo TodoRepo) GetTodos() []Todo {
	var todos []Todo
	repo.DB.Find(&todos)
	return todos
}

// GetTodo returns todo with ID id
func (repo TodoRepo) GetTodo(id int) Todo {
	var todo Todo
	repo.DB.First(&todo, id)
	return todo
}

// TodoRepo repository
type TodoRepo struct {
	DB *gorm.DB
}
