// internal/contact/service.go
package contact

import (
	"contact-api/internal/models"
	"errors"
)

type Service interface {
	CreateContact(contact *models.Contact) error
	GetContactWithFamily(id uint) (*models.Contact, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateContact(contact *models.Contact) error {
	// Business Rule: Validation
	if contact.Name == "" {
		return errors.New("contact name is required")
	}

	// Business Rule: Ensure it's not its own parent
	if contact.ParentID != nil && contact.ID != 0 && *contact.ParentID == contact.ID {
		return errors.New("a contact cannot be its own parent")
	}

	return s.repo.Create(contact)
}

func (s *service) GetContactWithFamily(id uint) (*models.Contact, error) {
	// This would call a repo method that uses GORM's .Preload("Children")
	if id == 0 {
		return nil, errors.New("invalid contact ID")
	}

	contact, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("contact not found")
	}

	return contact, nil
}
