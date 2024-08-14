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
	"golang.org/x/exp/slog"
)

type WearableData struct {
	WearableData *mongo.Collection
	LifeStyle    *mongo.Collection
}

func NewWearableData(mongoDb *mongo.Database) *WearableData {
	return &WearableData{
		WearableData: mongoDb.Collection("wearable_data"),
		LifeStyle:    mongoDb.Collection("lifestyles"),
	}
}

func (m *WearableData) Create(req *pb.WearableDataCreate) (*pb.Void, error) {
	data := &pb.WearableData{
		Id:                uuid.New().String(),
		UserId:            req.UserId,
		DeviceType:        req.DeviceType,
		StepCount:         req.StepCount,
		CaloriesBurned:    req.CaloriesBurned,
		DistanceTraveled:  req.DistanceTraveled,
		HeartRate:         req.HeartRate,
		SleepDuration:     req.SleepDuration,
		WorkoutType:       req.WorkoutType,
		Temperature:       req.Temperature,
		RecordedTimestamp: req.RecordedTimestamp,
		CreatedAt:         time.Now().Format(layout),
		UpdatedAt:         time.Now().Format(layout),
		DeletedAt:         0,
	}

	_, err := m.WearableData.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{Key: "recordeddate", Value: req.RecordedTimestamp[10:]},
		{Key: "deletedat", Value: 0},
	}

	res := &pb.LifeStyle{}

	err = m.LifeStyle.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		slog.Error("error finding life style for date: %v", err)
	}

	lf, err := m.CalculateDailyAverages(req.UserId, req.RecordedTimestamp)
	if err != nil {
		slog.Error("error calculating daily averages: %v", err)
	}

	res.StressLevel = int32(CalculateStressLevel(lf.HeartRate, float64(lf.SleepDuration), lf.StepCount))

	filter = bson.D{
		{Key: "recordeddate", Value: req.RecordedTimestamp[10:]},
		{Key: "deletedat", Value: 0},
	}

	updateFields := bson.D{
		{Key: "updatedat", Value: time.Now().Format(layout)},
	}

	if lf.HeartRate != 0 {
		updateFields = append(updateFields, bson.E{Key: "heartrate", Value: lf.HeartRate})
	}
	if lf.Temperature != 0 {
		updateFields = append(updateFields, bson.E{Key: "temperature", Value: lf.Temperature})
	}
	if lf.SleepDuration != 0 {
		updateFields = append(updateFields, bson.E{Key: "sleepduration", Value: lf.SleepDuration})
	}
	if res.StressLevel != 0 {
		updateFields = append(updateFields, bson.E{Key: "stresslevel", Value: res.StressLevel})
    }

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err = m.LifeStyle.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Error("error updating life style: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *WearableData) Update(req *pb.WearableDataUpdate) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	updateFields := bson.D{
		{Key: "updatedat", Value: time.Now().Format(layout)},
	}

	if req.DeviceType != "" {
		updateFields = append(updateFields, bson.E{Key: "devicetype", Value: req.DeviceType})
	}
	if req.StepCount != 0 {
		updateFields = append(updateFields, bson.E{Key: "stepcount", Value: req.StepCount})
	}
	if req.CaloriesBurned != 0 {
		updateFields = append(updateFields, bson.E{Key: "caloriesburned", Value: req.CaloriesBurned})
	}
	if req.DistanceTraveled != 0 {
		updateFields = append(updateFields, bson.E{Key: "distancetraveled", Value: req.DistanceTraveled})
	}
	if req.HeartRate != 0 {
		updateFields = append(updateFields, bson.E{Key: "heartrate", Value: req.HeartRate})
	}
	if req.SleepDuration != 0 {
		updateFields = append(updateFields, bson.E{Key: "sleepduration", Value: req.SleepDuration})
	}
	if req.WorkoutType != "" {
		updateFields = append(updateFields, bson.E{Key: "workouttype", Value: req.WorkoutType})
	}
	if req.RecordedTimestamp != "" {
		updateFields = append(updateFields, bson.E{Key: "recordedtimestamp", Value: req.RecordedTimestamp})
	}
	if req.Temperature != 0 {
		updateFields = append(updateFields, bson.E{Key: "temperature", Value: req.Temperature})
	}

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err := m.WearableData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}


	return &pb.Void{}, nil
}

func (m *WearableData) Delete(req *pb.GetById) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deletedat", Value: time.Now().Unix()},
		}},
	}

	_, err := m.WearableData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *WearableData) Get(req *pb.GetById) (*pb.WearableDataRes, error) {
	res := &pb.WearableDataRes{}

	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	err := m.WearableData.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *WearableData) List(req *pb.GetAllWearableDataReq) (*pb.GetAllWearableDataRes, error) {
	res := &pb.GetAllWearableDataRes{}

	filter := bson.D{{Key: "deletedat", Value: 0}}
	if req.UserId != "" {
		filter = append(filter, bson.E{Key: "userid", Value: req.UserId})
	}

	if req.RecordedTimestamp != "" {
		recordDate, err := time.Parse("2006-01-02", req.RecordedTimestamp)
		if err != nil {
			return nil, fmt.Errorf("invalid record date format: %v", err)
		}

		filter = append(filter, bson.E{Key: "recordedtimestamp", Value: bson.D{{Key: "$lt", Value: recordDate}}})
	}

	findOptions := options.Find()
	if req.Filter != nil {
		findOptions.SetLimit(int64(req.Filter.Limit))
		findOptions.SetSkip(int64(req.Filter.Offset))
	}

	cur, err := m.WearableData.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var doc pb.WearableDataRes
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		res.WearableData = append(res.WearableData, &doc)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *WearableData) GetRealTimeMonitoringData(req *pb.GetById) (*pb.WearableDataRes, error) {
	filter := bson.D{
		{Key: "userid", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "recordedtimestamp", Value: -1}})
	var wearableData pb.WearableDataRes
	err := m.WearableData.FindOne(context.Background(), filter, opts).Decode(&wearableData)
	if err != nil {
		return nil, err
	}

	return &wearableData, nil
}