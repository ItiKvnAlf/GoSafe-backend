package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Contact struct
type Contact struct {
	gorm.Model

	ID    uuid.UUID `gorm:"not null;unique_index" json:"id"`
	Name  string    `gorm:"not null;unique_index" json:"name"`
	Email string    `gorm:"not null;unique_index" json:"email"`
	Phone string    `gorm:"not null;unique_index" json:"phone"`
}
