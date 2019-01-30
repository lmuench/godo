package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/godo/config"
	"github.com/lmuench/godo/orm"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	router := httprouter.New()
	db := orm.InitPG()
	defer db.Close()

	config.InitRoutes(router, db)

	n.UseFunc(contentTypeJSON)
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":1323", n))
}

func contentTypeJSON(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Add("Content-Type", "application/json")
	next(w, r)
}
