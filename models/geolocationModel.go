package models

import (
	"github.com/google/uuid"
)

// Geolocation struct
type Geolocation struct {
	ID            uuid.UUID `gorm:"primaryKey;" json:"id"`
	CurrentPoint  string    `gorm:"not null" json:"current_point"`
	TravelRouteID uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignkey:TravelRouteID" json:"travel_route_id"`
}
