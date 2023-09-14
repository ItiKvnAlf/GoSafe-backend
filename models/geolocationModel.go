package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Geolocation struct
type Geolocation struct {
	gorm.Model

	ID         uuid.UUID `gorm:"not null;unique_index" json:"id"`
	Date       time.Time `gorm:"not null;unique_index" json:"date"`
	StartPoint string    `gorm:"not null;unique_index" json:"start_point"`
	EndPoint   string    `gorm:"not null;unique_index" json:"end_point"`
	Notes      string    `gorm:"not null;unique_index" json:"notes"`
}
