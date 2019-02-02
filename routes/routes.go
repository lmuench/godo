package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/api"
)

// InitRoutes registers routes
func InitRoutes(router *httprouter.Router, db *gorm.DB) {
	todoAPI := api.NewTodoAPI(db)
	userAPI := api.NewUserAPI(db)
	router.GET("/todos", todoAPI.GetTodos)
	router.GET("/todos/:id", todoAPI.GetTodo)
	router.POST("/sign-up", userAPI.SignUp)
	router.POST("/sign-in", userAPI.SignIn)
}
