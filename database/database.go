package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Sortren/event-log/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func MakeAutoMigrations() {
	err := DBConn.AutoMigrate(new(models.Event))

	if err != nil {
		log.Fatal("Can't auto migrate the database")
	}
	log.Print("Auto migrations went correctly")
}

func InitDatabaseConn() {
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	SSL_MODE := os.Getenv("SSL_MODE")
	TIMEZONE := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		POSTGRES_HOST,
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_DB, POSTGRES_PORT,
		SSL_MODE,
		TIMEZONE)

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed, check the provided credentials")
	}

	log.Print("Database connection successful")
}
