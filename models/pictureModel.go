package models

import (
	"github.com/google/uuid"
)

// PictureModel struct
type Picture struct {
	ID            uuid.UUID `gorm:"primaryKey;" json:"id"`
	Image         string    `gorm:"not null;" json:"image"`
	TravelRouteID uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignkey:TravelRouteID" json:"travel_route_id"`
}
