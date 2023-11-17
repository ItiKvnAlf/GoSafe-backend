package models

import (
	"github.com/google/uuid"
)

// Message struct
type Message struct {
	ID             uuid.UUID `gorm:"not null;unique_index" json:"id"`
	DefaultMessage string    `gorm:"not null" json:"default_message"`
}
