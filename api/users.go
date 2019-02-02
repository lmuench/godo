package api

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/models"
	"golang.org/x/crypto/bcrypt"
)

// SignUp ...
func (api UserAPI) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&_user)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if api.repo.UsernameTaken(_user.Username) {
		http.Error(w, "Username already taken", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(_user.Password), 8)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: _user.Username,
		Password: string(hashedPassword),
	}
	if err := api.repo.CreateUser(user); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// SignIn ...
func (api UserAPI) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// UserAPI controller
type UserAPI struct {
	repo models.UserRepo
}

// NewUserAPI constructor
func NewUserAPI(db *gorm.DB) UserAPI {
	repo := models.UserRepo{DB: db}
	api := UserAPI{repo}
	return api
}
