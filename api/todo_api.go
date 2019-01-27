package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	todo "github.com/lmuench/godo/models"
)

// GET returns all todos
func (h Handler) GET(ctx echo.Context) error {
	todos := h.repo.GetAll()
	return ctx.JSON(http.StatusOK, todos)
}

// Handler - Todo Controller
type Handler struct {
	repo *todo.Repo
}

// Register routes
func Register(e *echo.Echo, db *gorm.DB) {
	repo := todo.Repo{DB: db}
	h := Handler{&repo}

	e.GET("/todos", h.GET)
}
