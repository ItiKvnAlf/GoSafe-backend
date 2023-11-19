package models

import (
	"github.com/google/uuid"
)

// Contact struct
type Contact struct {
	ContactID uuid.UUID `gorm:"primaryKey;unique_index" json:"id"`
	Name      string    `gorm:"not null;unique_index" json:"name"`
	Email     string    `gorm:"not null;unique_index" json:"email"`
	Phone     string    `gorm:"not null;unique_index" json:"phone"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
}
