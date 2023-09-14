package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"not null;unique_index" json:"id"`
	Name       string    `gorm:"not null;unique_index" json:"name"`
	Email      string    `gorm:"not null;unique_index" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	Phone      string    `gorm:"not null;unique_index" json:"phone"`
	Address    string    `gorm:"not null" json:"address"`
	ProfilePic string    `gorm:"not null" json:"profile_pic"`
	Rut        string    `gorm:"not null;unique_index" json:"rut"`
	Picture    []Picture
}
