package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MedicalRecord struct {
	MedicalRecord *mongo.Collection
}

func NewMedicalRecord(mongoDb *mongo.Database) *MedicalRecord {
	return &MedicalRecord{
		MedicalRecord: mongoDb.Collection("medical_records"),
	}
}

var (
	layout = "2006-01-02 15:04:05"
)

func (m *MedicalRecord) Create(req *pb.MedicalRecordCreate) (*pb.Void, error) {

	record := &pb.MedicalRecord{
		Id:          uuid.New().String(),
		UserId:      req.UserId,
		RecordType:  req.RecordType,
		RecordDate:  req.RecordDate, // Make sure it's in time.Time format
		Description: req.Description,
		DoctorId:    req.DoctorId,
		Attachments: req.Attachments, // Make sure this is a slice of strings
		CreatedAt:   time.Now().Format(layout),
		UpdatedAt:   time.Now().Format(layout),
		DeletedAt:   0,
	}

	_, err := m.MedicalRecord.InsertOne(context.Background(), record)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *MedicalRecord) Update(req *pb.MedicalRecordUpdate) (*pb.Void, error) {
	res := &pb.Void{}

	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "doctorid", Value: req.DoctorId},
		{Key: "deletedat", Value: 0},
	}

	updateFields := bson.D{
		{Key: "updatedat", Value: time.Now().Format("2006-01-02 15:04:05")},
	}

	if req.RecordDate != "" && req.RecordDate != "string" {
		updateFields = append(updateFields, bson.E{Key: "recorddate", Value: req.RecordDate})
	}

	if req.RecordType != "" && req.RecordType != "string" {
		updateFields = append(updateFields, bson.E{Key: "recordtype", Value: req.RecordType})
	}

	if req.Description != "" && req.Description != "string" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: req.Description})
	}

	if req.Attachments != "" && req.Attachments != "string" {
		updateFields = append(updateFields, bson.E{Key: "attachments", Value: req.Attachments})
	}

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err := m.MedicalRecord.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MedicalRecord) Delete(req *pb.GetById) (*pb.Void, error) {
	res := &pb.Void{}

	filter := bson.D{{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deletedat", Value: time.Now().Unix()},
		}},
	}

	_, err := m.MedicalRecord.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MedicalRecord) Get(req *pb.GetById) (*pb.MedicalRecordRes, error) {
	res := &pb.MedicalRecordRes{}

	filter := bson.D{{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	err := m.MedicalRecord.FindOne(context.TODO(), filter).Decode(res)
	if err == mongo.ErrNoDocuments {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MedicalRecord) List(req *pb.GetAllMedicalRecords) (*pb.GetAllMedicalRecordsRes, error) {
	res := &pb.GetAllMedicalRecordsRes{}

	filter := bson.D{{Key: "deletedat", Value: 0}}
	if req.DoctorId != "" {
		filter = append(filter, bson.E{Key: "doctorid", Value: req.DoctorId})
	}

	if req.RecordType != "" {
		filter = append(filter, bson.E{Key: "recordtype", Value: req.RecordType})
    }

	if req.RecordDate != "" {
		recordDate, err := time.Parse("2006-01-02", req.RecordDate)
		if err != nil {
			return nil, fmt.Errorf("invalid record date format: %v", err)
		}

		filter = append(filter, bson.E{Key: "recorddate", Value: bson.D{{Key: "$lt", Value: recordDate}}})
	}

	if req.UserId!= "" {
        filter = append(filter, bson.E{Key: "userid", Value: req.UserId})
    }

	findOptions := options.Find()
	if req.Filter != nil {
		findOptions.SetLimit(int64(req.Filter.Limit))
		findOptions.SetSkip(int64(req.Filter.Offset))
	}

	cur, err := m.MedicalRecord.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var doc pb.MedicalRecordRes
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		res.MedicalRecords = append(res.MedicalRecords, &doc)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
