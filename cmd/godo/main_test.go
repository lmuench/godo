package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/lmuench/godo/internal/app/godo/middleware"
	"github.com/lmuench/godo/internal/app/godo/routes"
	"github.com/lmuench/godo/internal/app/godo/routes/handlers"
	"github.com/lmuench/godo/internal/pkg/services"
	"github.com/lmuench/godo/internal/platform/cache"
	"github.com/lmuench/godo/internal/platform/orm"
	"github.com/urfave/negroni"
)

var n *negroni.Negroni
var db *gorm.DB

func TestMain(m *testing.M) {
	c := cache.GetRedisConn()
	db = orm.InitPostgresTest()
	defer db.Close()

	router := routes.InitRoutes(
		c,
		handlers.NewTodoAPI(services.Todos{DB: db}, c),
		handlers.NewUserAPI(services.Users{DB: db}, c),
		handlers.NewOAuth2API(c),
	)

	n = negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("web/static")),
	)
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
		res := executeRequest(req)
		got := res.Code
		if got != c.expected {
			t.Errorf("Case: %s\nExpected response code %d. Got %d.\n", c.creds, c.expected, got)
		}
	}
}
