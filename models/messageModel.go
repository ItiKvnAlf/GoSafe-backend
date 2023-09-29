package models

import (
	"github.com/google/uuid"
)

// Message struct
type Message struct {
	ID             uuid.UUID `gorm:"not null;unique_index" json:"id"`
	UserID         uuid.UUID `gorm:"not null" json:"user_id"`
	TravelRouteID  uuid.UUID `gorm:"not null" json:"travel_route_id"`
	GeolocationID  uuid.UUID `gorm:"not null" json:"geolocation_id"`
	DefaultMessage string    `gorm:"not null" json:"default_message"`
	LastPicture    uuid.UUID `gorm:"not null" json:"last_picture"`

	User        User        `gorm:"foreignkey:UserID"`
	TravelRoute TravelRoute `gorm:"foreignkey:TravelRouteID"`
	Geolocation Geolocation `gorm:"foreignkey:GeolocationID"`
}
