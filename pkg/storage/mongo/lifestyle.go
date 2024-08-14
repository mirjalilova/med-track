package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Lifestyle struct {
	Lifestyle *mongo.Collection
}

func NewLifestyle(mongoDb *mongo.Database) *Lifestyle {
	return &Lifestyle{
		Lifestyle: mongoDb.Collection("lifestyles"),
	}
}

func (m *Lifestyle) Create(req *pb.LifestyleCreate) (*pb.Void, error) {
	lifestyle := &pb.LifeStyle{
		Id:            uuid.New().String(),
		UserId:        req.UserId,
		DataType:      req.DataType,
		DataValue:     req.DataValue,
		HeartRate:     req.HeartRate,
		BloodPressure: req.BloodPressure,
		Weight:        req.Weight,
		Height:        req.Height,
		Temperature:   req.Temperature,
		StressLevel:   req.StressLevel,
		SleepDuration: req.SleepDuration,
		RecordedDate:  req.RecordedDate,
		CreatedAt:     time.Now().Format(layout),
		UpdatedAt:     time.Now().Format(layout),
		DeletedAt:     0,
	}

	_, err := m.Lifestyle.InsertOne(context.Background(), lifestyle)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *Lifestyle) Update(req *pb.LifestyleUpdate) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	updateFields := bson.D{
		{Key: "updatedat", Value: time.Now().Format(layout)},
	}

	if req.DataType != "" {
		updateFields = append(updateFields, bson.E{Key: "datatype", Value: req.DataType})
	}
	if req.DataValue != "" {
		updateFields = append(updateFields, bson.E{Key: "datavalue", Value: req.DataValue})
	}
	if req.HeartRate != 0 {
		updateFields = append(updateFields, bson.E{Key: "heartrate", Value: req.HeartRate})
	}
	if req.BloodPressure != "" {
		updateFields = append(updateFields, bson.E{Key: "bloodpressure", Value: req.BloodPressure})
	}
	if req.Weight != 0 {
		updateFields = append(updateFields, bson.E{Key: "weight", Value: req.Weight})
	}
	if req.Height != 0 {
		updateFields = append(updateFields, bson.E{Key: "height", Value: req.Height})
	}
	if req.Temperature != 0 {
		updateFields = append(updateFields, bson.E{Key: "temperature", Value: req.Temperature})
	}
	if req.StressLevel != 0 {
		updateFields = append(updateFields, bson.E{Key: "stresslevel", Value: req.StressLevel})
	}
	if req.SleepDuration != 0 {
		updateFields = append(updateFields, bson.E{Key: "sleepduration", Value: req.SleepDuration})
	}

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err := m.Lifestyle.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *Lifestyle) Delete(req *pb.GetById) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deletedat", Value: time.Now().Unix()},
		}},
	}

	_, err := m.Lifestyle.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *Lifestyle) Get(req *pb.GetById) (*pb.LifestyleRes, error) {
	res := &pb.LifestyleRes{}

	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	err := m.Lifestyle.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *Lifestyle) List(req *pb.GetAllLifestyleReq) (*pb.GetAllLifestyleRes, error) {
	res := &pb.GetAllLifestyleRes{}

	filter := bson.D{{Key: "deletedat", Value: 0}}
	if req.UserId != "" {
		filter = append(filter, bson.E{Key: "userid", Value: req.UserId})
	}
	if req.DataType != "" {
		filter = append(filter, bson.E{Key: "datatype", Value: req.DataType})
	}

	findOptions := options.Find()
	if req.Filter != nil {
		findOptions.SetLimit(int64(req.Filter.Limit))
		findOptions.SetSkip(int64(req.Filter.Offset))
	}

	cur, err := m.Lifestyle.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var doc pb.LifestyleRes
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		res.Lifestyles = append(res.Lifestyles, &doc)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
