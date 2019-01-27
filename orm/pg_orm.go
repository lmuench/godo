package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For configuration
	"github.com/lmuench/godo/models"
)

// InitPG to Postgres with models
func InitPG() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=5432 dbname=godo_dev user=postgres password=postgres",
	)
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Todo{})
	return db
}