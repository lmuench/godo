package orm

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For configuration
	"github.com/lmuench/godo/models"
	"github.com/qor/admin"
)

// InitDevPG automigrates models and returns DB connection pointer
func InitDevPG() (*gorm.DB, *admin.Admin) {
	conf := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("GODO_DEV_DB_HOST"),
		os.Getenv("GODO_DEV_DB_PORT"),
		os.Getenv("GODO_DEV_DB_DBNAME"),
		os.Getenv("GODO_DEV_DB_USER"),
		os.Getenv("GODO_DEV_DB_PASSWORD"),
	)
	db, err := gorm.Open("postgres", conf)
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.User{})

	adm := admin.New(&admin.AdminConfig{DB: db})
	adm.AddResource(&models.Todo{})
	adm.AddResource(&models.User{})

	return db, adm
}

// InitEmptyTestPG drops tables, automigrates models and returns DB connection pointer
func InitEmptyTestPG() *gorm.DB {
	conf := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("GODO_TEST_DB_HOST"),
		os.Getenv("GODO_TEST_DB_PORT"),
		os.Getenv("GODO_TEST_DB_DBNAME"),
		os.Getenv("GODO_TEST_DB_USER"),
		os.Getenv("GODO_TEST_DB_PASSWORD"),
	)
	db, err := gorm.Open("postgres", conf)
	if err != nil {
		panic("failed to connect to database")
	}

	db.DropTableIfExists(
		&models.Todo{},
		&models.User{},
	)
	db.AutoMigrate(
		&models.Todo{},
		&models.User{},
	)
	return db
}
