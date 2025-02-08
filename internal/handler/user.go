package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rivaldi-fsociety/notification-service/proto"
	"github.com/rivaldi-fsociety/user-service/internal/service"
)

type UserHandler struct {
	Notification *service.NotificationClient
}

func NewUserHandler(notification *service.NotificationClient) *UserHandler {
	return &UserHandler{Notification: notification}
}

type RegisterUserRequest struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUser API
func (h *UserHandler) RegisterUser(c echo.Context) error {
	req := new(RegisterUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Simulate storing user (you can add DB later)
	log.Printf("✅ User registered: %s (%s)", req.UserID, req.Email)

	// Call Notification Service via gRPC
	_, err := h.Notification.Client.SendNotification(context.Background(), &proto.SendNotificationRequest{
		UserId:  req.UserID,
		Message: "Welcome! Your account has been created.",
	})
	if err != nil {
		log.Printf("❌ Failed to send notification: %v", err)
	} else {
		log.Println("📩 Notification sent successfully!")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}
