package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/cache"
	"github.com/lmuench/godo/middleware"
	"github.com/lmuench/godo/orm"
	"github.com/lmuench/godo/routes"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	router := httprouter.New()
	db, adm := orm.InitDevPG()
	defer db.Close()
	c := cache.GetRedisConn()

	routes.InitRoutes(router, db, c)

	n.UseFunc(middleware.CORS)

	mux := http.NewServeMux()
	mux.Handle("/", router)
	adm.MountTo("/admin", mux)
	n.UseHandler(mux)

	log.Println("Server listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", n))
}
