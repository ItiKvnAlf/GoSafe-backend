package models

import (
	"github.com/google/uuid"
)

// Message struct
type Message struct {
	MessageID      uuid.UUID `gorm:"primaryKey;unique_index" json:"id_message"`
	DefaultMessage string    `gorm:"not null" json:"default_message"`

	TravelRoute TravelRoute `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TravelRouteID"`
}
