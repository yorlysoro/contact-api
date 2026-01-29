// internal/models/contact.go
package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique"`
	Phone    string    `json:"phone" gorm:"not null"`
	UserID   uint      `json:"user_id"`   // Owner of the contact
	ParentID *uint     `json:"parent_id"` // For related/parent contacts
	Children []Contact `gorm:"foreignkey:ParentID"`
}
