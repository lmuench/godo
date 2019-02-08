package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/internal/pkg/services"
)

// GetTodos returns all todos
func (api TodoAPI) GetTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := api.service.GetTodos()
	respondWithJSON(w, todos)
}

// GetTodo returns all todos
func (api TodoAPI) GetTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	todos := api.service.GetTodo(id)
	respondWithJSON(w, todos)
}

// RespondWithJSON parses v as JSON and responds with "OK" and v or "Bad Request"
func respondWithJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// TodoAPI controller
type TodoAPI struct {
	service services.Todos
	cache   redis.Conn
}

// NewTodoAPI constructor
func NewTodoAPI(service services.Todos, cache redis.Conn) TodoAPI {
	api := TodoAPI{service, cache}
	return api
}
