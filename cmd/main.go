package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rivaldi-fsociety/user-service/internal/handler"
	"github.com/rivaldi-fsociety/user-service/internal/service"
)

func main() {
	e := echo.New()

	// Initialize gRPC Notification Client
	notificationClient := service.NewNotificationClient()

	// Initialize User Handler
	userHandler := handler.NewUserHandler(notificationClient)

	// Define Routes
	e.POST("/register", userHandler.RegisterUser)

	// Start HTTP Server
	log.Println("🚀 User Service running on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("❌ Failed to start HTTP server: %v", err)
	}
}
