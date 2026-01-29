// internal/auth/middleware_test.go
package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Should return 401 if no header is present", func(t *testing.T) {
		r := gin.New()
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) { c.Status(200) })

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
