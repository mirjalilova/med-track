package service

import (
	"context"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/pkg/storage/mongo"
	"github.com/mirjalilova/med-track/pkg/websocket"
	"golang.org/x/exp/slog"
)

type NotificationService struct {
	storage mongo.Storage
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(storage *mongo.Storage) *NotificationService {
	return &NotificationService{
		storage: *storage,
	}
}

func (s *NotificationService) Create(ctx context.Context, req *pb.Notification) (*pb.Void, error) {
	go func() {
		message := []byte(req.Message)
		websocket.BroadcastMessage(message)
	}()

	res, err := s.storage.Notification().Create(req)
	if err != nil {
		slog.Error("Error sending notification: %v", err)
		return nil, err
	}

	slog.Info("Successfully send notification")
	return res, nil
}

func (s *NotificationService) Get(ctx context.Context, req *pb.GetById) (*pb.Notification, error) {
	res, err := s.storage.Notification().Get(req)
	if err != nil {
		slog.Error("Error getting notification: %v", err)
		return nil, err
	}

	slog.Info("Get notification")
	return res, nil
}

func (s *NotificationService) List(ctx context.Context, req *pb.GetAllNotificationReq) (*pb.GetAllNotificationRes, error) {
	res, err := s.storage.Notification().List(req)
	if err != nil {
		slog.Error("Error listing notifications: %v", err)
		return nil, err
	}

	slog.Info("List notifications")
	return res, nil
}
