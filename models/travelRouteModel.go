package models

import (
	"time"

	"github.com/google/uuid"
)

// TravelRoute struct
type TravelRoute struct {
	ID            uuid.UUID `gorm:"not null;unique_index" json:"id"`
	StartPoint    string    `gorm:"not null" json:"start_point"`
	EndPoint      string    `gorm:"not null" json:"end_point"`
	Date          time.Time `gorm:"not null" json:"date"`
	UserID        uuid.UUID `gorm:"not null" json:"user_id"`
	PictureID     uuid.UUID `gorm:"not null" json:"last_picture_id"`
	MessageID     uuid.UUID `gorm:"foreignkey:MessageID"`
	GeolocationID uuid.UUID `gorm:"foreignkey:GeolocationID"`

	User        User        `gorm:"foreignkey:UserID"`
	Message     Message     `gorm:"foreignkey:MessageID"`
	Geolocation Geolocation `gorm:"foreignkey:GeolocationID"`
	Pictures    []Picture   `gorm:"many2many;jointable_foreignkey:TravelRouteID"` // Add a many-to-many relationship with Picture
}
