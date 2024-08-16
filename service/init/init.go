package service

import (
	st "github.com/mirjalilova/med-track/pkg/storage/mongo"
	"github.com/mirjalilova/med-track/service"
	"golang.org/x/exp/slog"
)

type Service struct {
	MedicalRecordS        service.MedicalRecordI
	GeneticDataS          service.GeneticDataI
	LifeStyleS            service.LifeStyleI
	WearableDataS         service.WearableDataI
	HealthRecommendationS service.HealthRecommendationI
	NotificationS         service.NotificationI
}

func ConnectDB() (*st.Storage, error) {
	storage, err := st.ConnectMongo()
	if err != nil {
		slog.Error("can't connect to mongo: %v", err)
		return nil, err
	}
	slog.Info("Connected to MongoDB")
	return storage, nil
}
