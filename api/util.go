package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

// HandleUserVerification returns username of logged-in user
// or replies with error status and returns error
func HandleUserVerification(w http.ResponseWriter, r *http.Request, cache redis.Conn) (string, error) {
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

// RespondWithJSON parses v as JSON and responds with "OK" and v or "Bad Request"
func RespondWithJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
