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

	if len(_user.Username) < 3 {
		http.Error(w, "Username must be at least 3 characters long", http.StatusBadRequest)
		return
	}

	if len(_user.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
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
	err = api.repo.CreateUser(user)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
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
