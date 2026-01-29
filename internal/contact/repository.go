// internal/contact/repository.go
package contact

import (
	"gorm.io/gorm"
)

// Contact model used by the repository (local copy to avoid importing external module)
type Contact struct {
	ID       uint
	Children []Contact
}

type Repository interface {
	Create(contact *Contact) error
	FindByID(id uint) (*Contact, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(contact *Contact) error {
	return r.db.Create(contact).Error
}

func (r *repository) FindByID(id uint) (*Contact, error) {
	var contact Contact
	// Preload("Children") allows us to see related contacts in one query
	err := r.db.Preload("Children").First(&contact, id).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}
