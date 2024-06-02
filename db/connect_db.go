package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		log.Fatal("Connecting to database failed", err)
	} else {
		fmt.Println("Connected to DB successfully")
	}
}
