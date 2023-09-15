package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Geolocation struct
type Geolocation struct {
	gorm.Model

	ID           uuid.UUID `gorm:"not null;unique_index" json:"id"`
	CurrentPoint string    `gorm:"not null" json:"current_pont"` // Change type
}
