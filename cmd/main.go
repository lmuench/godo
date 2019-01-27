package main

import (
	todo_api "github.com/lmuench/godo/api"
	pg_connector "github.com/lmuench/godo/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db := pg_connector.ConnectAndMigrate()
	defer db.Close()

	echo := echo.New()

	// register routes
	todo_api.Register(echo, db)

	echo.Logger.Fatal(echo.Start(":1323"))
}
