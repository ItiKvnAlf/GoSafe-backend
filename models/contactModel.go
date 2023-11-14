package models

import (
	"github.com/google/uuid"
)

// Contact struct
type Contact struct {
	ID     uuid.UUID `gorm:"not null;unique_index" json:"id"`
	UserID uuid.UUID `gorm:"not null" json:"user_id"`
	Name   string    `gorm:"not null;unique_index" json:"name"`
	Email  string    `gorm:"not null;unique_index" json:"email"`
	Phone  string    `gorm:"not null;unique_index" json:"phone"`
}
