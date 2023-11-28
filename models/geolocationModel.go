package models

import (
	"github.com/google/uuid"
)

// Geolocation struct
type Geolocation struct {
	ID            uuid.UUID `gorm:"primaryKey;" json:"id"`
	CurrentPoint  float64   `gorm:"not null" json:"current_point"`
	TravelRouteID uuid.UUID `gorm:"not null" json:"travel_route_id"`

	TravelRoute Travel_route `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignkey:TravelRouteID"`
}
