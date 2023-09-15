package models

import (
	"github.com/google/uuid"
)

// Geolocation struct
type Geolocation struct {
	ID           uuid.UUID `gorm:"not null;unique_index" json:"id"`
	CurrentPoint string    `gorm:"not null" json:"current_pont"` // Change type
}
