package models

import (
	"github.com/google/uuid"
)

// Geolocation struct
type Geolocation struct {
	GeolocationID uuid.UUID `gorm:"primaryKey;unique_index" json:"geolocation_id"`
	CurrentPoint  string    `gorm:"not null" json:"current_point"` // Change type

	TravelRoute TravelRoute `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TravelRouteID"`
}
