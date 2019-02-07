package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/models"
	uuid "github.com/satori/go.uuid"
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

	token := uuid.Must(uuid.NewV4()).String()
	_, err = api.cache.Do("SETEX", token, "120", user.Username)
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
	})
}

// Refresh ...
func (api UserAPI) Refresh(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	oldToken, err := GetToken(w, r)
	if err != nil {
		return
	}
	username, err := GetUsername(w, api.cache, oldToken)
	if err != nil {
		return
	}

	newToken := uuid.Must(uuid.NewV4()).String()
	_, err = api.cache.Do("SETEX", newToken, "120", username)
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	_, err = api.cache.Do("DEL", oldToken)
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    newToken,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
	})
}

// UserAPI controller
type UserAPI struct {
	repo  models.UserRepo
	cache redis.Conn
}

// NewUserAPI constructor
func NewUserAPI(repo models.UserRepo, cache redis.Conn) UserAPI {
	api := UserAPI{repo, cache}
	return api
}
