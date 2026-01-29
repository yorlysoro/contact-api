// internal/contact/service_test.go
package contact

import (
	"contact-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock type for the Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(contact *models.Contact) error {
	args := m.Called(contact)
	return args.Error(0)
}

func (m *MockRepository) FindByID(id uint) (*models.Contact, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Contact), args.Error(1)
}

func TestService_CreateContact(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	t.Run("Success - Create Contact", func(t *testing.T) {
		contact := &models.Contact{Name: "Jane Doe", Phone: "555-0101"}

		// Setup expectation
		mockRepo.On("Create", contact).Return(nil)

		err := svc.CreateContact(contact)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
