package database

import (
	"fmt"
	"log"
	"os"

	"github.com/AaronDennis07/electrum/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database ")
	}
	fmt.Println("DB connected")

	log.Println("Connected")
	// db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	err = db.AutoMigrate(&models.Department{}, &models.Student{}, &models.Course{}, &models.Session{}, &models.Enrollment{}, &models.Admin{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database migrated")
	DB = Dbinstance{
		Db: db,
	}
}
