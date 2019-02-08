package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/internal/app/godo/middleware"
	"github.com/lmuench/godo/internal/app/godo/routes"
	"github.com/lmuench/godo/internal/app/godo/routes/handlers"
	"github.com/lmuench/godo/internal/pkg/services"
	"github.com/lmuench/godo/internal/platform/cache"
	"github.com/lmuench/godo/internal/platform/orm"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	router := httprouter.New()
	db, adm := orm.InitPostgresDev()
	defer db.Close()
	c := cache.GetRedisConn()

	todoAPI := handlers.NewTodoAPI(services.Todos{DB: db}, c)
	userAPI := handlers.NewUserAPI(services.Users{DB: db}, c)
	oauth2API := handlers.NewOAuth2API(c)
	routes.InitRoutes(router, c, todoAPI, userAPI, oauth2API)

	n.UseFunc(middleware.CORS)

	mux := http.NewServeMux()
	mux.Handle("/", router)
	adm.MountTo("/admin", mux)
	n.UseHandler(mux)

	log.Println("Server listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", n))
}
