package api

import (
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/models"
)

// GetTodos returns all todos
func (api TodoAPI) GetTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if _, err := GetToken(w, r); err != nil {
		return
	}
	todos := api.repo.GetTodos()
	RespondWithJSON(w, todos)
}

// GetTodo returns all todos
func (api TodoAPI) GetTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	todos := api.repo.GetTodo(id)
	RespondWithJSON(w, todos)
}

// TodoAPI controller
type TodoAPI struct {
	repo  models.TodoRepo
	cache redis.Conn
}

// NewTodoAPI constructor
func NewTodoAPI(db *gorm.DB, cache redis.Conn) TodoAPI {
	repo := models.TodoRepo{DB: db}
	api := TodoAPI{repo, cache}
	return api
}
