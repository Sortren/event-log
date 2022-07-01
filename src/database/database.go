package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabaseConn() {
	// TODO: move credentials to .env, change host to service name before deployment
	dsn := "host=localhost user=sortren password=sortren123 dbname=main port=5432 sslmode=disable TimeZone=Europe/Warsaw"

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")
	}
	fmt.Println("Database connection successful")

}
