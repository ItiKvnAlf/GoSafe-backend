package models

import (
	"github.com/google/uuid"
)

// PictureModel struct
type Picture struct {
	PictureID uuid.UUID `gorm:"primaryKey;unique_index" json:"picture_id"`
	Image     string    `gorm:"not null;unique_index" json:"image"`

	TravelRoute TravelRoute `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TravelRouteID"`
}
