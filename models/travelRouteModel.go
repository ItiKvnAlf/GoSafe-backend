package models

import (
	"time"

	"github.com/google/uuid"
)

// TravelRoute struct
type TravelRoute struct {
	ID         uuid.UUID `gorm:"not null;unique_index" json:"id"`
	StartPoint string    `gorm:"not null" json:"start_point"`
	EndPoint   string    `gorm:"not null" json:"end_point"`
	Date       time.Time `gorm:"not null" json:"date"`
	UserID     uuid.UUID `gorm:"not null" json:"user_id"`
	Pictures   []Picture `gorm:"foreignkey:TravelRouteID"`

	User User `gorm:"foreignkey:UserID"`
}
