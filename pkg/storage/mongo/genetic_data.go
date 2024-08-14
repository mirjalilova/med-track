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

type GeneticData struct {
	GeneticData *mongo.Collection
}

func NewGeneticData(mongoDb *mongo.Database) *GeneticData {
	return &GeneticData{
		GeneticData: mongoDb.Collection("genetic_data"),
	}
}

func (m *GeneticData) Create(req *pb.GeneticDataCreate) (*pb.Void, error) {
	data := &pb.GeneticData{
		Id:           uuid.New().String(),
		UserId:       req.UserId,
		DataType:     req.DataType,
		DataValue:    req.DataValue,
		AnalysisDate: req.AnalysisDate,
		DoctorId:     req.DoctorId,
		CreatedAt:    time.Now().Format(layout),
		UpdatedAt:    time.Now().Format(layout),
		DeletedAt:    0,
	}

	_, err := m.GeneticData.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *GeneticData) Update(req *pb.GeneticDataUpdate) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "doctorid", Value: req.DoctorId},
		{Key: "deletedat", Value: 0},
	}

	updateFields := bson.D{
		{Key: "datatype", Value: req.DataType},
		{Key: "datavalue", Value: req.DataValue},
		{Key: "analysisdate", Value: req.AnalysisDate},
	}

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err := m.GeneticData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *GeneticData) Delete(req *pb.GetById) (*pb.Void, error) {
	filter := bson.D{{Key: "id", Value: req.Id}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deletedat", Value: time.Now().Unix()},
		}},
	}

	_, err := m.GeneticData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *GeneticData) Get(req *pb.GetById) (*pb.GeneticDataRes, error) {
	res := &pb.GeneticDataRes{}

	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	err := m.GeneticData.FindOne(context.TODO(), filter).Decode(res)
	if err == mongo.ErrNoDocuments {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *GeneticData) List(req *pb.GetAllGeneticDateReq) (*pb.GetAllGeneticDateRes, error) {
	res := &pb.GetAllGeneticDateRes{}

	filter := bson.D{
		{Key: "deletedat", Value: 0},
	}

	if req.UserId != "" {
		filter = append(filter, bson.E{Key: "userid", Value: req.UserId})
	}

	if req.DoctorId != "" {
		filter = append(filter, bson.E{Key: "doctorid", Value: req.DoctorId})
	}

	if req.AnalysisDate != "" {
		analysisDate, err := time.Parse("2006-01-02", req.AnalysisDate)
		if err != nil {
			return nil, fmt.Errorf("invalid analysis date format: %v", err)
		}

		filter = append(filter, bson.E{Key: "analysisdate", Value: bson.D{{Key: "$lt", Value: analysisDate}}})
	}

	findOptions := options.Find()
	if req.Filter != nil {
		findOptions.SetLimit(int64(req.Filter.Limit))
		findOptions.SetSkip(int64(req.Filter.Offset))
	}

	cur, err := m.GeneticData.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var doc pb.GeneticDataRes
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		res.GeneticDatas = append(res.GeneticDatas, &doc)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	res.Count = int32(len(res.GeneticDatas))
	return res, nil
}
