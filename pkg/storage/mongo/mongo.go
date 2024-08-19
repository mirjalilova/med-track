package mongo

import (
	"context"

	"github.com/mirjalilova/med-track/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB                    *mongo.Database
	MedicalRecordS        storage.MedicalRecordI
	GeneticDataS          storage.GeneticDataI
	LifeStyleS            storage.LifeStyleI
	WearableDataS         storage.WearableDataI
	HealthRecommendationS storage.HealthRecommendationI
	NotificationS         storage.NotificationI
}

func ConnectMongo() (*Storage, error) {
	// time.Sleep(time.Second*10)
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017").
	SetAuth(options.Credential{Username: "mongo", Password: "feruza1727"})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// time.Sleep(time.Second*10)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database("medmon")

	return &Storage{
		DB:             db,
		MedicalRecordS: NewMedicalRecord(db),
	}, nil
}

func (s *Storage) MedicalRecord() storage.MedicalRecordI {
	if s.MedicalRecordS == nil {
		s.MedicalRecordS = NewMedicalRecord(s.DB)
	}
	return s.MedicalRecordS
}

func (s *Storage) GeneticData() storage.GeneticDataI {
	if s.GeneticDataS == nil {
		s.GeneticDataS = NewGeneticData(s.DB)
	}
	return s.GeneticDataS
}

func (s *Storage) LifeStyle() storage.LifeStyleI {
	if s.LifeStyleS == nil {
		s.LifeStyleS = NewLifestyle(s.DB)
	}
	return s.LifeStyleS
}

func (s *Storage) WearableData() storage.WearableDataI {
	if s.WearableDataS == nil {
		s.WearableDataS = NewWearableData(s.DB)
	}
	return s.WearableDataS
}

func (s *Storage) HealthRecommendation() storage.HealthRecommendationI {
	if s.HealthRecommendationS == nil {
		s.HealthRecommendationS = NewRecommendation(s.DB)
	}
	return s.HealthRecommendationS
}

func (s *Storage) Notification() storage.NotificationI {
	if s.NotificationS == nil {
        s.NotificationS = NewNotification(s.DB)
    }
    return s.NotificationS
}
