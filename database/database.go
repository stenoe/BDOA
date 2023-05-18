package database

import (
	"fmt"
	"log"
	"os"

	"github.com/stenoe/go-database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=Europe/Tallinn",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// This happens when there is and error!
	if err != nil {
		log.Fatal("Can't connect to database \n", err)
		os.Exit(2)
	}

	// This happens when there is no error
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// tell the database how its table look alike
	log.Println("running mogrations")
	db.AutoMigrate(&models.Fact{})

	// Store the opened database into its variable
	DB = Dbinstance{
		Db: db,
	}
}

// Fira Code ligatures  ! =  !=  : = :=
