package models

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Todo Model
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetAll returns all todos
func (r *Repo) GetAll(c echo.Context) error {
	var todos []Todo
	r.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

// Repo - Todo Repository
type Repo struct {
	DB *gorm.DB
}
