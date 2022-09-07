package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/amirhossein-programmer/cmd/web/pkg/models"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("MyApp.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: \n", err)
		os.Exit(2)
	}
	log.Println("Connecting database: ", db.Name())
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	// migrations
	db.AutoMigrate(&models.User{}, &models.Product{})

	Database = DbInstance{DB: db}
}
