package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/internal/pkg/services"
	"github.com/lmuench/godo/internal/pkg/services/models"
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

	err = api.service.SignUp(_user)
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

	user, err := api.service.GetUser(_user.Username)
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
	oldToken, err := getToken(w, r)
	if err != nil {
		return
	}
	username, err := getUsername(w, api.cache, oldToken)
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

// getToken returns session token of current client
// or replies with error status and returns error
func getToken(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "You need to login first", http.StatusUnauthorized)
			return "", err
		}
		http.Error(w, "Oops, something went wrong!", http.StatusBadRequest)
		return "", err
	}
	token := cookie.Value
	return token, nil
}

// getUsername returns username mapped to session token
// or replies with error status and returns error
func getUsername(w http.ResponseWriter, cache redis.Conn, token string) (string, error) {
	reply, err := cache.Do("GET", token)
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return "", err
	}
	username, err := redis.String(reply, err)
	if err != nil {
		http.Error(w, "You need to login again!", http.StatusUnauthorized)
		return "", err
	}
	return username, nil
}

// UserAPI controller
type UserAPI struct {
	service services.Users
	cache   redis.Conn
}

// NewUserAPI constructor
func NewUserAPI(service services.Users, cache redis.Conn) UserAPI {
	api := UserAPI{service, cache}
	return api
}
