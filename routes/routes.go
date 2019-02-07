package routes

import (
	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/api"
	"github.com/lmuench/godo/oauth"
)

// InitRoutes registers routes
func InitRoutes(
	router *httprouter.Router,
	cache redis.Conn,
	todoAPI api.TodoAPI,
	userAPI api.UserAPI,
	oauthAPI oauth.API,
) {
	router.GET("/todos", todoAPI.GetTodos)
	router.GET("/todos/:id", todoAPI.GetTodo)
	router.POST("/sign-up", userAPI.SignUp)
	router.POST("/sign-in", userAPI.SignIn)
	router.POST("/refresh", userAPI.Refresh)
	router.GET("/oauth/redirect", oauthAPI.Redirect)
}
