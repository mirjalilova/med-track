package service

import (
	"context"

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
    return res, nil
}

func (s *LifeStyleService) Update(ctx context.Context, req *pb.LifestyleUpdate) (*pb.Void, error) {
	res, err := s.storage.LifeStyle().Update(req)
    if err != nil {
		slog.Error("Error updating lifestyle: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.LifeStyle().Delete(req)
    if err != nil {
		slog.Error("Error deleting lifestyle: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) Get(ctx context.Context, req *pb.GetById) (*pb.LifestyleRes, error) {
	res, err := s.storage.LifeStyle().Get(req)
    if err != nil {
		slog.Error("Error getting lifestyle: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) List(ctx context.Context, req *pb.GetAllLifestyleReq) (*pb.GetAllLifestyleRes, error) {
	res, err := s.storage.LifeStyle().List(req)
    if err!= nil {
		slog.Error("Error listing lifestyles: %v", err)
        return nil, err
    }

    return res, nil
}