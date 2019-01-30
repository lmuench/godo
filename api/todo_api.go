package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lmuench/godo/models"
)

// GetAllTodos returns all todos
func (api TodoAPI) GetAllTodos(ctx echo.Context) error {
	todos := api.repo.GetAllTodos()
	return ctx.JSON(http.StatusOK, todos)
}

// TodoAPI controller
type TodoAPI struct {
	repo *models.TodoRepo
}

// RegisterTodo registers todo routes
func RegisterTodo(echo *echo.Echo, db *gorm.DB) {
	repo := models.TodoRepo{DB: db}
	api := TodoAPI{&repo}

	echo.GET("/todos", api.GetAllTodos)
}
