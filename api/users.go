package api

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/models"
	"golang.org/x/crypto/bcrypt"
)

// SignUp creates a new user from the "username" and "password" params
func (api UserAPI) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var _user models.User
	err := json.NewDecoder(r.Body).Decode(&_user)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	err = api.repo.SignUp(_user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// SignIn ...
func (api UserAPI) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var _user models.User
	err := json.NewDecoder(r.Body).Decode(&_user)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	user, err := api.repo.GetUser(_user.Username)
	if err != nil {
		http.Error(w, "Username doesn't exist", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(_user.Password))
	if err != nil {
		http.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	//TODO
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
