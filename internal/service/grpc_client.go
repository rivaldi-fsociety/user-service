package service

import (
	"log"

	"github.com/rivaldi-fsociety/notification-service/proto"
	"google.golang.org/grpc"
)

type NotificationClient struct {
	Client proto.NotificationServiceClient
}

func NewNotificationClient() *NotificationClient {
	// Connect to Notification Service gRPC Server
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Could not connect to Notification Service: %v", err)
	}

	client := proto.NewNotificationServiceClient(conn)
	return &NotificationClient{Client: client}
}
