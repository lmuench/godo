package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/api"
	"github.com/lmuench/godo/cache"
	"github.com/lmuench/godo/middleware"
	"github.com/lmuench/godo/models"
	"github.com/lmuench/godo/oauth"
	"github.com/lmuench/godo/orm"
	"github.com/lmuench/godo/routes"
	"github.com/urfave/negroni"
)

var n *negroni.Negroni
var db *gorm.DB

func TestMain(m *testing.M) {
	n = negroni.Classic()
	router := httprouter.New()
	db = orm.InitEmptyTestPG()
	defer db.Close()
	c := cache.GetRedisConn()

	todoAPI := api.NewTodoAPI(models.TodoRepo{DB: db}, c)
	userAPI := api.NewUserAPI(models.UserRepo{DB: db}, c)
	oauthAPI := oauth.NewAPI(c)
	routes.InitRoutes(router, c, todoAPI, userAPI, oauthAPI)

	n.UseFunc(middleware.CORS)
	n.UseHandler(router)

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	n.ServeHTTP(rec, req)
	return rec
}

func TestSignUp(t *testing.T) {
	cases := []struct {
		creds    string
		expected int
	}{
		{`{ "username": "john", "password": "hello123" }`, http.StatusCreated},
		{`{ "username": "sara", "password": "12345678" }`, http.StatusCreated},
		{`{ "username": "john", "password": "aehuihiuh" }`, http.StatusBadRequest},
		{`{ "username": "foofoo", "password": "abc" }`, http.StatusBadRequest},
	}

	for _, c := range cases {
		reqBody := []byte(c.creds)
		req, _ := http.NewRequest("POST", "/sign-up", bytes.NewBuffer(reqBody))
		got := executeRequest(req).Code
		if got != c.expected {
			t.Errorf("Case: %s\nExpected response code %d. Got %d.\n", c.creds, c.expected, got)
		}
	}
}
