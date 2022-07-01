package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	Description string `json:"description"`
	Type        string `json:"type"`
}
