package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
 dsn := os.Getenv("DATABASE_URL")
	log.Println("DATABASE_URL:", dsn)
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
		return
	}
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    DB = db
}
