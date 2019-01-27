package main

import (
	router "github.com/lmuench/godo/conf"
	pg_connector "github.com/lmuench/godo/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db := pg_connector.Init()
	defer db.Close()

	e := echo.New()

	router.Init(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
