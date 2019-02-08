package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lmuench/godo/internal/pkg/services/types"
)

// GetTodos returns all todos
func (s Todos) GetTodos() []types.Todo {
	var todos []types.Todo
	s.DB.Find(&todos)
	return todos
}

// GetTodo returns todo with provided ID
func (s Todos) GetTodo(id int) types.Todo {
	var todo types.Todo
	s.DB.First(&todo, id)
	return todo
}

// Todos ...
type Todos struct {
	DB *gorm.DB
}
