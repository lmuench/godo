package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lmuench/godo/internal/pkg/services/models"
)

// GetTodos returns all todos
func (s Todos) GetTodos() []models.Todo {
	var todos []models.Todo
	s.DB.Find(&todos)
	return todos
}

// GetTodo returns todo with provided ID
func (s Todos) GetTodo(id int) models.Todo {
	var todo models.Todo
	s.DB.First(&todo, id)
	return todo
}

// Todos ...
type Todos struct {
	DB *gorm.DB
}
