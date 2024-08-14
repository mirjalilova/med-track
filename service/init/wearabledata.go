package service

import (
	"context"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/pkg/storage/mongo"
	"golang.org/x/exp/slog"
)

type WearableDataService struct {
	storage mongo.Storage
	pb.UnimplementedWearableDataServiceServer
}

func NewWearableDataService(storage *mongo.Storage) *WearableDataService {
	return &WearableDataService{
		storage: *storage,
	}
}

func (s *WearableDataService) Create(ctx context.Context, req *pb.WearableDataCreate) (*pb.Void, error) {
	res, err := s.storage.WearableData().Create(req)
    if err != nil {
		slog.Error("Error creating wearable data: %v", err)
        return nil, err
    }

    slog.Info("Successfully created wearable data")
    return res, nil
}

func (s *WearableDataService) Update(ctx context.Context, req *pb.WearableDataUpdate) (*pb.Void, error) {
	res, err := s.storage.WearableData().Update(req)
    if err != nil {
		slog.Error("Error updating wearable data: %v", err)
        return nil, err
    }

    slog.Info("Successfully updated wearable data")
    return res, nil
}

func (s *WearableDataService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.WearableData().Delete(req)
    if err != nil {
		slog.Error("Error deleting wearable data: %v", err)
        return nil, err
    }

    slog.Info("Successfully deleted wearable data")
    return res, nil
}

func (s *WearableDataService) Get(ctx context.Context, req *pb.GetById) (*pb.WearableDataRes, error) {
	res, err := s.storage.WearableData().Get(req)
    if err != nil {
		slog.Error("Error getting wearable data: %v", err)
        return nil, err
    }

    slog.Info("Successfully retrieved wearable data")
    return res, nil
}

func (s *WearableDataService) List(ctx context.Context, req *pb.GetAllWearableDataReq) (*pb.GetAllWearableDataRes, error) {
	res, err := s.storage.WearableData().List(req)
    if err!= nil {
		slog.Error("Error listing wearable data: %v", err)
        return nil, err
    }

    slog.Info("Successfully listed wearable data")
    return res, nil
}

func (s *WearableDataService) GetRealTimeMonitoringData(ctx context.Context, req *pb.GetById) (*pb.WearableDataRes, error) {
    res, err := s.storage.WearableData().GetRealTimeMonitoringData(req)
    if err!= nil {
		slog.Error("Error: getting real-time monitoring data: %v", err)
        return nil, err
    }

    slog.Info("Successfully retrieved real-time monitoring data")
    return res, nil
}