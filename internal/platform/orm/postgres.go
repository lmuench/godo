package orm

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For configuration
	"github.com/lmuench/godo/internal/pkg/services/types"
	"github.com/qor/admin"
)

// InitPostgresDev automigrates gorm models and returns DB connection pointer
func InitPostgresDev() (*gorm.DB, *admin.Admin) {
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

	db.AutoMigrate(&types.Todo{})
	db.AutoMigrate(&types.User{})

	adm := admin.New(&admin.AdminConfig{DB: db})
	adm.AddResource(&types.Todo{})
	adm.AddResource(&types.User{})

	return db, adm
}

// InitPostgresTest drops tables, automigrates gorm models and returns DB connection pointer
func InitPostgresTest() *gorm.DB {
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
		&types.Todo{},
		&types.User{},
	)
	db.AutoMigrate(
		&types.Todo{},
		&types.User{},
	)
	return db
}
