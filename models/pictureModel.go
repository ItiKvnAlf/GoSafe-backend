package models

import (
	"github.com/google/uuid"
)

// PictureModel struct
type Picture struct {
	ID    uuid.UUID `gorm:"not null;unique_index" json:"id"`
	Image string    `gorm:"not null;unique_index" json:"image"`
}
