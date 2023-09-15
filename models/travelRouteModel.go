package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TravelRoute struct
type TravelRoute struct {
	gorm.Model

	ID         uuid.UUID `gorm:"not null;unique_index" json:"id"`
	StartPoint string    `gorm:"not null" json:"start_point"`
	EndPoint   string    `gorm:"not null" json:"end_point"`
	Date       time.Time `gorm:"not null" json:"date"`
	Pictures   []Picture
}
