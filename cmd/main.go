package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/middleware"
	"github.com/lmuench/godo/orm"
	"github.com/lmuench/godo/routes"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	router := httprouter.New()
	db := orm.InitPG()
	defer db.Close()

	routes.InitRoutes(router, db)

	n.UseFunc(middleware.ContentTypeJSON)
	n.UseFunc(middleware.CORS)
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":1323", n))
}
