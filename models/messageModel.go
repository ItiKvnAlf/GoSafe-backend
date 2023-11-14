package models

import (
	"github.com/google/uuid"
)

// Message struct
type Message struct {
	MessageID      uuid.UUID `gorm:"not null;unique_index" json:"message_id"`
	DefaultMessage string    `gorm:"not null" json:"default_message"`
}
