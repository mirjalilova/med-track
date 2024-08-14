package service

import (
	"context"
	"time"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/pkg/storage/mongo"
	"golang.org/x/exp/slog"
)

type LifeStyleService struct {
	storage mongo.Storage
	pb.UnimplementedLifestyleServiceServer
}

func NewLifeStyleService(storage *mongo.Storage) *LifeStyleService {
	return &LifeStyleService{
		storage: *storage,
	}
}

func (s *LifeStyleService) Create(ctx context.Context, req *pb.LifestyleCreate) (*pb.Void, error) {
	res, err := s.storage.LifeStyle().Create(req)
	if err != nil {
		slog.Error("Error creating lifestyle: %v", err)
		return nil, err
	}

    slog.Info("Successfully created lifestyle")
	return res, nil
}

func (s *LifeStyleService) Update(ctx context.Context, req *pb.LifestyleUpdate) (*pb.Void, error) {
	res, err := s.storage.LifeStyle().Update(req)
	if err != nil {
		slog.Error("Error updating lifestyle: %v", err)
		return nil, err
	}

    slog.Info("Successfully updated lifestyle")
	return res, nil
}

func (s *LifeStyleService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.LifeStyle().Delete(req)
	if err != nil {
		slog.Error("Error deleting lifestyle: %v", err)
		return nil, err
	}

    slog.Info("Delete lifestyle")
	return res, nil
}

func (s *LifeStyleService) Get(ctx context.Context, req *pb.GetById) (*pb.LifestyleRes, error) {
	res, err := s.storage.LifeStyle().Get(req)
	if err != nil {
		slog.Error("Error getting lifestyle: %v", err)
		return nil, err
	}

    slog.Info("Get lifestyle")
	return res, nil
}

func (s *LifeStyleService) List(ctx context.Context, req *pb.GetAllLifestyleReq) (*pb.GetAllLifestyleRes, error) {
	res, err := s.storage.LifeStyle().List(req)
	if err != nil {
		slog.Error("Error listing lifestyles: %v", err)
		return nil, err
	}

    slog.Info("List lifestyles")
	return res, nil
}

func (s *LifeStyleService) GetWeeklySummary(ctx context.Context, req *pb.WeeklySummaryReq) (*pb.WeeklySummary, error) {
	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")
	
    res, err := s.storage.LifeStyle().GetWeeklySummary(&pb.WeeklySummaryReq{
        UserId: req.UserId,
        StartDate: startDate,
        EndDate: endDate,  // 7 days ago from now. You can adjust this as needed.
    })
    if err!= nil {
        slog.Error("Error getting weekly summary: %v", err)
        return nil, err
    }

    slog.Info("Get weekly summary")
    return res, nil
}
