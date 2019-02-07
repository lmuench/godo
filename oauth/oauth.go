package oauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
)

// Redirect handler function
func (api API) Redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		respondWithLoginError(w, err)
		return
	}
	code := r.FormValue("code")
	if len(code) == 0 {
		respondWithLoginError(w, err)
		return
	}

	url := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		os.Getenv("GODO_DEV_GITHUB_OAUTH_CLIENT_ID"),
		os.Getenv("GODO_DEV_GITHUB_OAUTH_CLIENT_SECRET"),
		code,
	)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		respondWithLoginError(w, err)
		return
	}
	req.Header.Set("accept", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		respondWithLoginError(w, err)
		return
	}
	defer res.Body.Close()

	var token AccessToken
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		respondWithLoginError(w, err)
		return
	}

	err = api.cacheTokenAndGitHubID(token)
	if err != nil {
		respondWithLoginError(w, err)
		return
	}

	w.Header().Set("Location", "/welcome.html?access_token="+token.Value)
	w.WriteHeader(http.StatusFound)
}

func (api API) cacheTokenAndGitHubID(token AccessToken) error {
	url := "https://api.github.com/user"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "token "+token.Value)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var id GitHubID
	err = json.NewDecoder(res.Body).Decode(&id)
	if err != nil {
		return err
	}

	api.cache.Do("SET", token.Value, id.Value)
	return nil
}

// AccessToken contains OAuth access token
type AccessToken struct {
	Value string `json:"access_token"`
}

// GitHubID contains GitHub ID
type GitHubID struct {
	Value int `json:"id"`
}

func respondWithLoginError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(
		w,
		"Oops, something went wrong. Please try logging in again",
		http.StatusBadRequest,
	)
}

// API receiver
type API struct {
	cache redis.Conn
}

// NewAPI constructor
func NewAPI(cache redis.Conn) API {
	o := API{cache}
	return o
}
