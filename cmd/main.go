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
// cmd/main.go
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yorlysoro/contact-api/internal/auth"
	"github.com/yorlysoro/contact-api/internal/contact"
	"github.com/yorlysoro/contact-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite" // Using SQLite for simplicity; swap for Postgres/MySQL easily
	"gorm.io/gorm"
)

func main() {

	// Load .env at the start
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	// 1. Initialize Database
	// In production, use environment variables for the connection string
	db, err := gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Run Auto-Migrations
	// This creates the tables based on your structs
	err = db.AutoMigrate(&models.Contact{}, &models.User{}) // Add &models.User{} when implemented
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
	setupRoutes(r, contactHandler, db)

	// 6. Start Server
	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(r *gin.Engine, ch *contact.Handler, db *gorm.DB) {
	authHandler := &auth.AuthHandler{DB: db}
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

		// Public Routes
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}
}
