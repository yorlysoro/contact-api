// internal/auth/service_test.go
package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "mysecurepassword"
	hashed, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEqual(t, password, hashed)

	// Verify
	match := CheckPasswordHash(password, hashed)
	assert.True(t, match)
}
