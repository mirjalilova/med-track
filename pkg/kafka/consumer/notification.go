package kafka

import (
	"context"
	"encoding/json"
	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	service "github.com/mirjalilova/med-track/service/init"
	"golang.org/x/exp/slog"
)

func CreateNotificationHandler(m *service.NotificationService) func(message []byte) {
	return func(message []byte) {
		var notification pb.Notification
		if err := json.Unmarshal(message, &notification); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := m.Create(context.Background(), &notification)
		if err != nil {
			slog.Error("failed to create notification via Kafka: %v", err)
			return
		}
		slog.Info("Create Notification")
	}
}
