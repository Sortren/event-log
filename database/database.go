package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDatabase() error {
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
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("database connection failed, err: %w", err)
	}

	log.Print("database connection successful")

	if err := migrate(); err != nil {
		return fmt.Errorf("can't run auto migrations based on models, err: %w", err)
	}
	log.Print("auto migrations went successfully")

	return nil
}

func GetConnection() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("database connection has not been initialized")
	}
	return db, nil
}
