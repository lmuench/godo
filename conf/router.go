package conf

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	todo "github.com/lmuench/godo/models"
)

// Init registers routes
func Init(e *echo.Echo, db *gorm.DB) {
	todo := todo.Repo{db}
	e.GET("/todos", todo.GetAll)
}
