package storage

import (
    pb "github.com/mirjalilova/med-track/genproto/health_sync"
)

type StorageI interface {
    MedicalRecord() MedicalRecordI
    GeneticData() GeneticDataI
    LifeStyle() LifeStyleI
    WearableData() WearableDataI
    HealthRecommendation() HealthRecommendationI
}

type MedicalRecordI interface {
    Create(req *pb.MedicalRecordCreate) (*pb.Void, error)
    Update(req *pb.MedicalRecordUpdate) (*pb.Void, error)
    Delete(req *pb.GetById) (*pb.Void, error)
    Get(req *pb.GetById) (*pb.MedicalRecordRes, error)
    List(req *pb.GetAllMedicalRecords) (*pb.GetAllMedicalRecordsRes, error)
}

type GeneticDataI interface {
    Create(req *pb.GeneticDataCreate) (*pb.Void, error)
    Update(req *pb.GeneticDataUpdate) (*pb.Void, error)
    Delete(req *pb.GetById) (*pb.Void, error)
    Get(req *pb.GetById) (*pb.GeneticDataRes, error)
    List(req *pb.GetAllGeneticDateReq) (*pb.GetAllGeneticDateRes, error)
}

type LifeStyleI interface {
    Create(req *pb.LifestyleCreate) (*pb.Void, error)
    Update(req *pb.LifestyleUpdate) (*pb.Void, error)
    Delete(req *pb.GetById) (*pb.Void, error)
    Get(req *pb.GetById) (*pb.LifestyleRes, error)
    List(req *pb.GetAllLifestyleReq) (*pb.GetAllLifestyleRes, error)
    GetWeeklySummary(req *pb.WeeklySummaryReq) (*pb.WeeklySummary, error)
}

type WearableDataI interface {
    Create(req *pb.WearableDataCreate) (*pb.Void, error)
    Update(req *pb.WearableDataUpdate) (*pb.Void, error)
    Delete(req *pb.GetById) (*pb.Void, error)
    Get(req *pb.GetById) (*pb.WearableDataRes, error)
    List(req *pb.GetAllWearableDataReq) (*pb.GetAllWearableDataRes, error)
    GetRealTimeMonitoringData(req *pb.GetById) (*pb.WearableDataRes, error)
}

type HealthRecommendationI interface {
    Create(req *pb.RecommendationCreate) (*pb.Void, error)
    Update(req *pb.RecommendationUpdate) (*pb.Void, error)
    Delete(req *pb.GetById) (*pb.Void, error)
    Get(req *pb.GetById) (*pb.RecommendationRes, error)
    List(req *pb.GetAllRecommendationReq) (*pb.GetAllRecommendation, error)
}

