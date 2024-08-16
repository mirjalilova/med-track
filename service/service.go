package service

import (
	"context"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
)

type ServiceI interface {
	MedicalRecord() MedicalRecordI
	GeneticData() GeneticDataI
	LifeStyle() LifeStyleI
	WearableData() WearableDataI
	HealthRecommendation() HealthRecommendationI
	Notification() NotificationI
}

type MedicalRecordI interface {
	Create(ctx context.Context, request *pb.MedicalRecordCreate) (*pb.Void, error)
	Update(ctx context.Context, request *pb.MedicalRecordUpdate) (*pb.Void, error)
	Delete(ctx context.Context, request *pb.GetById) (*pb.Void, error)
	Get(ctx context.Context, request *pb.GetById) (*pb.MedicalRecordRes, error)
	List(ctx context.Context, request *pb.GetAllMedicalRecords) (*pb.GetAllMedicalRecordsRes, error)
}

type GeneticDataI interface {
	Create(ctx context.Context, request *pb.GeneticDataCreate) (*pb.Void, error)
	Update(ctx context.Context, request *pb.GeneticDataUpdate) (*pb.Void, error)
	Delete(ctx context.Context, request *pb.GetById) (*pb.Void, error)
	Get(ctx context.Context, request *pb.GetById) (*pb.GeneticDataRes, error)
	List(ctx context.Context, request *pb.GetAllGeneticDateReq) (*pb.GetAllGeneticDateRes, error)
}

type LifeStyleI interface {
	Create(ctx context.Context, request *pb.LifestyleCreate) (*pb.Void, error)
	Update(ctx context.Context, request *pb.LifestyleUpdate) (*pb.Void, error)
	Delete(ctx context.Context, request *pb.GetById) (*pb.Void, error)
	Get(ctx context.Context, request *pb.GetById) (*pb.LifestyleRes, error)
	List(ctx context.Context, request *pb.GetAllLifestyleReq) (*pb.GetAllLifestyleRes, error)
	GetWeeklySummary(ctx context.Context, request *pb.WeeklySummaryReq) (*pb.WeeklySummary, error)
}

type WearableDataI interface {
	Create(ctx context.Context, request *pb.WearableDataCreate) (*pb.Void, error)
	Update(ctx context.Context, request *pb.WearableDataUpdate) (*pb.Void, error)
	Delete(ctx context.Context, request *pb.GetById) (*pb.Void, error)
	Get(ctx context.Context, request *pb.GetById) (*pb.WearableDataRes, error)
	List(ctx context.Context, request *pb.GetAllWearableDataReq) (*pb.GetAllWearableDataRes, error)
	GetRealTimeMonitoringData(ctx context.Context, request *pb.GetById) (*pb.WearableDataRes, error)
}

type HealthRecommendationI interface {
	Create(ctx context.Context, request *pb.RecommendationCreate) (*pb.Void, error)
	Update(ctx context.Context, request *pb.RecommendationUpdate) (*pb.Void, error)
	Delete(ctx context.Context, request *pb.GetById) (*pb.Void, error)
	Get(ctx context.Context, request *pb.GetById) (*pb.RecommendationRes, error)
	List(ctx context.Context, request *pb.GetAllRecommendationReq) (*pb.GetAllRecommendation, error)
}

type NotificationI interface {
	Create(ctx context.Context, request *pb.Notification) (*pb.Void, error)
	Get(ctx context.Context, request *pb.GetById) (*pb.Notification, error)
	List(ctx context.Context, request *pb.GetAllNotificationReq) (*pb.GetAllNotificationRes, error)
}
