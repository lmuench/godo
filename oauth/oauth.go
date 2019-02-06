package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// Redirect handler function
func (api API) Redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		respondWithLoginError(w)
		return
	}
	code := r.FormValue("code")
	if len(code) == 0 {
		respondWithLoginError(w)
		return
	}

	reqURL := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		os.Getenv("GODO_DEV_GITHUB_OAUTH_CLIENT_ID"),
		os.Getenv("GODO_DEV_GITHUB_OAUTH_CLIENT_SECRET"),
		code,
	)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		respondWithLoginError(w)
		return
	}
	req.Header.Set("accept", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		respondWithLoginError(w)
		return
	}
	defer res.Body.Close()

	var token AccessToken
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		respondWithLoginError(w)
		return
	}

	api.cache.Do("SET", token.Value, "1")

	w.Header().Set("Location", "/welcome.html?access_token="+token.Value)
	w.WriteHeader(http.StatusFound)
}

// AccessToken contains OAuth access token
type AccessToken struct {
	Value string `json:"access_token"`
}

func respondWithLoginError(w http.ResponseWriter) {
	http.Error(
		w,
		"Oops, something went wrong. Please try logging in again",
		http.StatusBadRequest,
	)
}

// API receiver
type API struct {
	db    *gorm.DB
	cache redis.Conn
}

// NewAPI constructor
func NewAPI(db *gorm.DB, cache redis.Conn) API {
	o := API{db, cache}
	return o
}
