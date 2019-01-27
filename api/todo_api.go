package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lmuench/godo/models"
)

// GetAll returns all todos
func (api TodoAPI) GetAll(ctx echo.Context) error {
	todos := api.repo.GetAll()
	return ctx.JSON(http.StatusOK, todos)
}

// TodoAPI - Todo Controller
type TodoAPI struct {
	repo *models.TodoRepo
}

// Register routes
func RegisterTodo(echo *echo.Echo, db *gorm.DB) {
	repo := models.TodoRepo{DB: db}
	api := TodoAPI{&repo}

	echo.GET("/todos", api.GetAll)
}
