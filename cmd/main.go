package main

import (
	"github.com/lmuench/godo/api"
	"github.com/lmuench/godo/orm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db := orm.InitPG()
	defer db.Close()

	echo := echo.New()

	// register routes
	api.RegisterTodo(echo, db)

	echo.Logger.Fatal(echo.Start(":1323"))
}
