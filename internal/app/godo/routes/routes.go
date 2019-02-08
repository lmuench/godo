package routes

import (
	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/internal/app/godo/routes/handlers"
)

// InitRoutes registers routes
func InitRoutes(
	router *httprouter.Router,
	cache redis.Conn,
	todoAPI handlers.TodoAPI,
	userAPI handlers.UserAPI,
	oauth2API handlers.OAuth2API,
) {
	router.GET("/todos", todoAPI.GetTodos)
	router.GET("/todos/:id", todoAPI.GetTodo)
	router.POST("/sign-up", userAPI.SignUp)
	router.POST("/sign-in", userAPI.SignIn)
	router.POST("/refresh", userAPI.Refresh)
	router.GET("/oauth/redirect", oauth2API.Redirect)
}
