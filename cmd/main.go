package main

import (
	todo_api "github.com/lmuench/godo/api"
	pg_connector "github.com/lmuench/godo/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db := pg_connector.Init()
	defer db.Close()

	e := echo.New()

	// register routes
	todo_api.Register(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
