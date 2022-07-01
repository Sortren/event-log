package database

import (
	"log"

	"github.com/Sortren/event-log/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func migrations() {
	DBConn.AutoMigrate(new(models.Event))
}

func InitDatabaseConn() {
	// TODO: move credentials to .env, change host to service name before deployment
	dsn := "host=localhost user=sortren password=sortren123 dbname=main port=5432 sslmode=disable TimeZone=Europe/Warsaw"

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}
	log.Print("Database connection successful")

	migrations()
}
