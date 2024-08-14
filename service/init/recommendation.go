package service

import (
	"context"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/pkg/storage/mongo"
	"golang.org/x/exp/slog"
)

type RecomendationService struct {
	storage mongo.Storage
	pb.UnimplementedRecommendationServiceServer
}

func NewRecomendationService(storage *mongo.Storage) *RecomendationService {
	return &RecomendationService{
		storage: *storage,
	}
}

func (s *RecomendationService) Create(ctx context.Context, req *pb.RecommendationCreate) (*pb.Void, error) {
	res, err := s.storage.HealthRecommendation().Create(req)
    if err != nil {
		slog.Error("Error creating recommendation: %v", err)
        return nil, err
    }

    slog.Info("Successfully created recommendation")
    return res, nil
}

func (s *RecomendationService) Update(ctx context.Context, req *pb.RecommendationUpdate) (*pb.Void, error) {
	res, err := s.storage.HealthRecommendation().Update(req)
    if err != nil {
		slog.Error("Error updating recommendation: %v", err)
        return nil, err
    }

    slog.Info("Successfully updated recommendation")
    return res, nil
}

func (s *RecomendationService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.HealthRecommendation().Delete(req)
    if err != nil {
		slog.Error("Error deleting recommendation: %v", err)
        return nil, err
    }

    slog.Info("Successfully deleted recommendation")
    return res, nil
}

func (s *RecomendationService) Get(ctx context.Context, req *pb.GetById) (*pb.RecommendationRes, error) {
	res, err := s.storage.HealthRecommendation().Get(req)
    if err != nil {
		slog.Error("Error getting recommendation: %v", err)
        return nil, err
    }

    slog.Info("Successfully retrieved recommendation")
    return res, nil
}

func (s *RecomendationService) List(ctx context.Context, req *pb.GetAllRecommendationReq) (*pb.GetAllRecommendation, error) {
	res, err := s.storage.HealthRecommendation().List(req)
    if err!= nil {
		slog.Error("Error listing recommendations: %v", err)
        return nil, err
    }

    slog.Info("Successfully listed recommendations")
    return res, nil
}