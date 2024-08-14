package service

import (
	"context"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/pkg/storage/mongo"
	"golang.org/x/exp/slog"
)

type GeneticDataService struct {
	storage mongo.Storage
	pb.UnimplementedGeneticDataServiceServer
}

func NewGeneticDataService(storage *mongo.Storage) *GeneticDataService {
	return &GeneticDataService{
		storage: *storage,
	}
}

func (s *GeneticDataService) Create(ctx context.Context, req *pb.GeneticDataCreate) (*pb.Void, error) {
	res, err := s.storage.GeneticData().Create(req)
    if err != nil {
		slog.Error("Error creating genetic data: %v", err)
        return nil, err
    }
    return res, nil
}


func (s *GeneticDataService) Update(ctx context.Context, req *pb.GeneticDataUpdate) (*pb.Void, error) {
	res, err := s.storage.GeneticData().Update(req)
    if err != nil {
		slog.Error("Error updating genetic data: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *GeneticDataService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.GeneticData().Delete(req)
    if err != nil {
		slog.Error("Error deleting genetic data: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *GeneticDataService) Get(ctx context.Context, req *pb.GetById) (*pb.GeneticDataRes, error) {
	res, err := s.storage.GeneticData().Get(req)
    if err != nil {
		slog.Error("Error getting genetic data: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *GeneticDataService) List(ctx context.Context, req *pb.GetAllGeneticDateReq) (*pb.GetAllGeneticDateRes, error) {
	res, err := s.storage.GeneticData().List(req)
    if err!= nil {
		slog.Error("Error listing genetic data: %v", err)
        return nil, err
    }

    return res, nil
}