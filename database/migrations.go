package database

import "github.com/Sortren/event-log/models"

func migrate() error {
	return db.AutoMigrate(
		&models.Event{},
	)
}
