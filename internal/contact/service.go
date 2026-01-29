/*
BSD 3-Clause License

Copyright (c) 2026, yorlysoro

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
	list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
	this list of conditions and the following disclaimer in the documentation
	and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
	contributors may be used to endorse or promote products derived from
	this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
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
