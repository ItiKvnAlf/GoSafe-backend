package models

import (
	"time"

	"github.com/google/uuid"
)

// TravelRoute struct
type Travel_route struct {
	ID         uuid.UUID `gorm:"primaryKey;" json:"id"`
	StartPoint string    `gorm:"not null" json:"start_point"`
	EndPoint   string    `gorm:"not null" json:"end_point"`
	Date       time.Time `gorm:"not null" json:"date"`
	UserID     uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignkey:UserID" json:"user_id"`

	Message     []Message     `gorm:"foreignkey:TravelRouteID"`
	Pictures    []Picture     `gorm:"foreignkey:TravelRouteID"`
	Geolocation []Geolocation `gorm:"foreignkey:TravelRouteID"`
}
