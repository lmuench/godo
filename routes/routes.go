package routes

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/api"
)

// InitRoutes registers routes
func InitRoutes(router *httprouter.Router, db *gorm.DB, cache redis.Conn) {
	todoAPI := api.NewTodoAPI(db, cache)
	userAPI := api.NewUserAPI(db, cache)
	router.GET("/todos", todoAPI.GetTodos)
	router.GET("/todos/:id", todoAPI.GetTodo)
	router.POST("/sign-up", userAPI.SignUp)
	router.POST("/sign-in", userAPI.SignIn)
	router.POST("/refresh", userAPI.Refresh)
}
