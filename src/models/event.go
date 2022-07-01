package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model

	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"`
}
