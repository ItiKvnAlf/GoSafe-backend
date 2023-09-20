package models

import (
	"github.com/google/uuid"
)

// User struct
type User struct {
	ID         uuid.UUID `gorm:"not null;unique" json:"id"`
	Name       string    `gorm:"not null;" json:"name"`
	Email      string    `gorm:"not null;unique" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	Phone      string    `gorm:"not null;unique" json:"phone"`
	Address    string    `gorm:"not null" json:"address"`
	ProfilePic string    `gorm:"not null" json:"profile_pic"`
	Rut        string    `gorm:"not null;unique" json:"rut"`
	Contacts   []Contact `gorm:"foreignkey:UserID"`
}
