package database

import (
	"fmt"
	"log"
	"os"

	"github.com/stenoe/server3/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var DB DBinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Tallinn",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Connect to the database
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)

	if err != nil {
		log.Fatal("Failed to connect to database!")
		os.Exit(2)
	}

	log.Println("Database connection successful.")

	// Migrate the schema
	db.AutoMigrate(&models.Measurement{})

	DB = DBinstance{
		Db: db,
	}
}
