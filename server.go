package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

// Todo model
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *Handler) getTodos(c echo.Context) error {
	var todos []Todo
	h.db.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

// Handler holds DB connection and request handling methods
type Handler struct {
	db *gorm.DB
}

func main() {
	db := migrateAndGetDB()
	defer db.Close()

	h := Handler{db}

	e := echo.New()
	e.GET("/todos", h.getTodos)
	e.Logger.Fatal(e.Start(":1323"))
}

func migrateAndGetDB() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=5432 dbname=godo_dev user=postgres password=postgres",
	)
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Todo{})
	return db
}
