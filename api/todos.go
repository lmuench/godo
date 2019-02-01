package api

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/models"
)

// GetTodos returns all todos
func (api TodoAPI) GetTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := api.repo.GetTodos()
	RespondWithJSON(w, todos)
}

// GetTodo returns all todos
func (api TodoAPI) GetTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	Handle400(w, err)
	todos := api.repo.GetTodo(id)
	RespondWithJSON(w, todos)
}

// TodoAPI controller
type TodoAPI struct {
	repo models.TodoRepo
}

// NewTodoAPI constructor
func NewTodoAPI(db *gorm.DB) TodoAPI {
	repo := models.TodoRepo{DB: db}
	api := TodoAPI{repo}
	return api
}
