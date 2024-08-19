package kafka

import (
	"context"
	"encoding/json"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/service/init"
	"golang.org/x/exp/slog"
)

func CreateMedicalRecordHandler(m *service.MedicalRecordService) func(message []byte) {
	return func(message []byte) {
		var med pb.MedicalRecordCreate
		if err := json.Unmarshal(message, &med); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := m.Create(context.Background(), &med)
		if err != nil {
			slog.Error("failed to create medical record via Kafka: %v", err)
			return
		}
		slog.Info("Create Medical Record")
	}
}

func UpdateMedicalRecorddHandler(m *service.MedicalRecordService) func(message []byte) {
	return func(message []byte) {
		var med pb.MedicalRecordUpdate
		if err := json.Unmarshal(message, &med); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := m.Update(context.Background(), &med)
		if err != nil {
			slog.Error("failed to edit password via Kafka: %v", err)
			return
		}
		slog.Info("Medocal Record updated")
	}
}