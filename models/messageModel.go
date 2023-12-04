package models

import (
	"github.com/google/uuid"
)

// Message struct
type Message struct {
	ID             uuid.UUID `gorm:"primaryKey;" json:"id"`
	DefaultMessage string    `gorm:"not null" json:"default_message"`
	TravelRouteID  uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignkey:TravelRouteID" json:"travel_route_id"`
}
