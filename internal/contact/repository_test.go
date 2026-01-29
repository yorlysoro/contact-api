// internal/contact/repository_test.go
package contact

import (
	"contact-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateContact(t *testing.T) {
	// Setup in-memory DB for testing
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Contact{})

	repo := NewRepository(db)

	t.Run("Should create a valid contact", func(t *testing.T) {
		contact := &models.Contact{Name: "John Doe", Phone: "123456"}
		err := repo.Create(contact)

		assert.NoError(t, err)
		assert.NotEqual(t, 0, contact.ID)
	})
}
