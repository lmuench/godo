package main

import (
	"log"
	"net/http"

	"github.com/lmuench/godo/internal/app/godo/middleware"
	"github.com/lmuench/godo/internal/app/godo/routes"
	"github.com/lmuench/godo/internal/app/godo/routes/handlers"
	"github.com/lmuench/godo/internal/pkg/services"
	"github.com/lmuench/godo/internal/platform/cache"
	"github.com/lmuench/godo/internal/platform/orm"
	"github.com/urfave/negroni"
)

func main() {
	c := cache.GetRedisConn()
	db, adm := orm.InitPostgresDev()
	defer db.Close()

	router := routes.InitRoutes(
		c,
		handlers.NewTodoAPI(services.Todos{DB: db}, c),
		handlers.NewUserAPI(services.Users{DB: db}, c),
		handlers.NewOAuth2API(c),
	)

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("web/static")),
	)
	n.UseFunc(middleware.CORS)

	mux := http.NewServeMux()
	mux.Handle("/", router)
	adm.MountTo("/admin", mux)
	n.UseHandler(mux)

	log.Println("Server listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", n))
}
