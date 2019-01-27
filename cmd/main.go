package main

import (
	"github.com/labstack/echo"
	"github.com/lmuench/godo/api"
	"github.com/lmuench/godo/orm"
)

func main() {
	db := orm.InitPG()
	defer db.Close()

	echo := echo.New()

	api.RegisterTodo(echo, db)

	echo.Logger.Fatal(echo.Start(":1323"))
}
