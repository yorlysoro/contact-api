// cmd/main.go
package main

import (
	"log"

	"contact-api/internal/auth"
	"contact-api/internal/contact"
	"contact-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite" // Using SQLite for simplicity; swap for Postgres/MySQL easily
	"gorm.io/gorm"
)

func main() {
	// 1. Initialize Database
	// In production, use environment variables for the connection string
	db, err := gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Run Auto-Migrations
	// This creates the tables based on your structs
	err = db.AutoMigrate(&models.Contact{}) // Add &models.User{} when implemented
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// 3. Dependency Injection
	// We wire up the layers from bottom to top
	contactRepo := contact.NewRepository(db)
	contactSvc := contact.NewService(contactRepo)
	contactHandler := contact.NewHandler(contactSvc)

	// 4. Initialize Gin Router
	r := gin.Default()

	// 5. Setup Routes
	setupRoutes(r, contactHandler)

	// 6. Start Server
	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(r *gin.Engine, ch *contact.Handler) {
	v1 := r.Group("/api/v1")
	{
		// Contact Routes - Protected by JWT Middleware
		contacts := v1.Group("/contacts")
		contacts.Use(auth.AuthMiddleware())
		{
			contacts.POST("/", ch.Create)
			contacts.GET("/:id", ch.GetByID)
			// Future: contacts.PUT("/:id", ch.Update)
			// Future: contacts.DELETE("/:id", ch.Delete)
		}

		// Public Routes (Example)
		// v1.POST("/register", userHandler.Register)
		// v1.POST("/login", userHandler.Login)
	}
}
