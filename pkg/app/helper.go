package app

import (
	"log/slog"

	service "github.com/mirjalilova/med-track/service/init"
	kafka "github.com/mirjalilova/med-track/pkg/kafka/consumer"
)

func RegisterConsumer(brokers []string, kcm *kafka.KafkaConsumerManager, medService *service.MedicalRecordService) {

	if err := kcm.RegisterConsumer(brokers, "cre-med", "med", kafka.CreateMedicalRecordHandler(medService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			slog.Warn("Consumer for topic 'cre-med' already exists")
		} else {
			slog.Error("Error registering consumer: %v", "err", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "upd-med", "med", kafka.UpdateMedicalRecorddHandler(medService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			slog.Warn("Consumer for topic 'upd-med' already exists")
		} else {
			slog.Error("Error registering consumer: %v", "err", err)
		}
	}

	// if err := kcm.RegisterConsumer(brokers, "upd-pass", "auth", kafka.UserEditPasswordHandler(medService)); err != nil {
	// 	if err == kafka.ErrConsumerAlreadyExists {
	// 		slog.Warn("Consumer for topic 'upd-pass' already exists")
	// 	} else {
	// 		slog.Error("Error registering consumer: %v", "err", err)
	// 	}
	// }

	// if err := kcm.RegisterConsumer(brokers, "upd-setting", "auth", kafka.UserEditSettingHandler(medService)); err != nil {
	// 	if err == kafka.ErrConsumerAlreadyExists {
	// 		slog.Warn("Consumer for topic 'upd-setting' already exists")
	// 	} else {
	// 		slog.Error("Error registering consumer: %v", "err", err)
	// 	}
	// }
}
