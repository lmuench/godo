package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/models"
)

// GetAllTodos returns all todos
func (api TodoAPI) GetAllTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := api.repo.GetAllTodos()
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
