package models

import (
	"github.com/google/uuid"
)

// PictureModel struct
type Picture struct {
	ID       uuid.UUID `gorm:"not null;unique_index" json:"id"`
	TravelID uuid.UUID `gorm:"not null" json:"travel_id"`
	Image    string    `gorm:"not null;unique_index" json:"image"`

	TravelRoute TravelRoute `gorm:"foreignkey:TravelID"`
}
