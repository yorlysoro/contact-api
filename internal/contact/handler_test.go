// internal/contact/handler_test.go
package contact

import (
	"contact-api/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService for testing handlers
type MockService struct {
	mock.Mock
}

func (m *MockService) CreateContact(c *models.Contact) error {
	return m.Called(c).Error(0)
}

func (m *MockService) GetContactWithFamily(id uint) (*models.Contact, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Contact), args.Error(1)
}

func TestHandler_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(MockService)
	h := NewHandler(mockSvc)

	t.Run("Success - Get Contact", func(t *testing.T) {
		mockSvc.On("GetContactWithFamily", uint(1)).Return(&models.Contact{Name: "Test"}, nil)

		r := gin.New()
		r.GET("/contacts/:id", h.GetByID)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/contacts/1", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockSvc.AssertExpectations(t)
	})
}
