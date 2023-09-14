package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PictureMode struct
type PictureMode struct {
	gorm.Model

	ID     uuid.UUID `gorm:"not null;unique_index" json:"id"`
	Image  string    `gorm:"not null;unique_index" json:"image"`
	UserID string    `gorm:"not null;unique_index" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID"`
}
