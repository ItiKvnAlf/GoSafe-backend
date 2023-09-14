package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Message struct
type Message struct {
	gorm.Model

	ID             uuid.UUID   `gorm:"not null;unique_index" json:"id"`
	User           User        `gorm:"foreignKey:UserID"`
	Picture        Picture     `gorm:"foreignKey:PictureID"`
	Geolocation    Geolocation `gorm:"foreignKey:GeolocationID"`
	Contact        Contact     `gorm:"foreignKey:ContactID"`
	DefaultMessage string      `gorm:"not null;unique_index" json:"default_message"`
}
