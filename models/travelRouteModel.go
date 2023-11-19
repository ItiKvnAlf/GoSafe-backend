package models

import (
	"time"

	"github.com/google/uuid"
)

// TravelRoute struct
type TravelRoute struct {
	TravelRouteIDID uuid.UUID `gorm:"not null;unique_index" json:"id"`
	StartPoint      string    `gorm:"not null" json:"start_point"`
	EndPoint        string    `gorm:"not null" json:"end_point"`
	Date            time.Time `gorm:"not null" json:"date"`

	User User `gorm:"foreignkey:UserID"`

	Message     []Message     `gorm:"foreignkey:TravelRouteID"`
	Geolocation []Geolocation `gorm:"foreignkey:TravelRouteID"`
	Picture     []Picture     `gorm:"foreignkey:TravelRouteID"`
}
