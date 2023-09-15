package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Message struct
type Message struct {
	gorm.Model

	ID             uuid.UUID `gorm:"not null;unique_index" json:"id"`
	User           User
	TravelRoute    TravelRoute
	Geolocation    Geolocation
	Contact        Contact
	DefaultMessage string `gorm:"not null" json:"default_message"`
	LastPicture    string `gorm:"not null" json:"lasg_picture"`
}
