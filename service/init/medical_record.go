package service

import (
	"context"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/pkg/storage/mongo"
	"golang.org/x/exp/slog"
)

type MedicalRecordService struct {
	storage mongo.Storage
	pb.UnimplementedMedicalRecordServiceServer
}

func NewMedicalRecordService(storage *mongo.Storage) *MedicalRecordService {
	return &MedicalRecordService{
		storage: *storage,
	}
}

func (s *MedicalRecordService) Create(ctx context.Context, req *pb.MedicalRecordCreate) (*pb.Void, error) {
	res, err := s.storage.MedicalRecord().Create(req)
	if err != nil {
        slog.Error("Error creating medical record: %v", err)
        return nil, err
    }
	return res, nil
}


func (s *MedicalRecordService) Update(ctx context.Context, req *pb.MedicalRecordUpdate) (*pb.Void, error) {
	res, err := s.storage.MedicalRecord().Update(req)
    if err != nil {
        slog.Error("Error updating medical record: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *MedicalRecordService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.MedicalRecord().Delete(req)
    if err != nil {
        slog.Error("Error deleting medical record: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *MedicalRecordService) Get(ctx context.Context, req *pb.GetById) (*pb.MedicalRecordRes, error) {
	res, err := s.storage.MedicalRecord().Get(req)
    if err != nil {
        slog.Error("Error getting medical record: %v", err)
        return nil, err
    }
    return res, nil
}

func (s *MedicalRecordService) List(ctx context.Context, req *pb.GetAllMedicalRecords) (*pb.GetAllMedicalRecordsRes, error) {
	res, err := s.storage.MedicalRecord().List(req)
	if err!= nil {
        slog.Error("Error listing medical records: %v", err)
        return nil, err
    }

    return res, nil
}