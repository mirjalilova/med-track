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
		StepCount:     req.StepCount,
		SleepDuration: req.SleepDuration,
		RecordedDate:  req.RecordedDate,
		CreatedAt:     time.Now().Format(layout),
		UpdatedAt:     time.Now().Format(layout),
		DeletedAt:     0,
	}

	lifestyle.StressLevel = int32(CalculateStressLevel(req.HeartRate, float64(req.SleepDuration), req.StepCount))

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

	req.StressLevel = int32(CalculateStressLevel(req.HeartRate, float64(req.SleepDuration), req.StepCount))

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

    filter := bson.M{
        "$or": []bson.M{
            {"id": req.Id},
            {"userid": req.Id},
        },
        "deletedat": 0,
    }

    err := m.Lifestyle.FindOne(context.TODO(), filter).Decode(res)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            fmt.Println("No documents found with the given filter.")
        } else {
            fmt.Printf("Error executing query: %v\n", err)
        }
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

func (m *Lifestyle) GetWeeklySummary(req *pb.WeeklySummaryReq) (*pb.WeeklySummary, error) {
	filter := bson.D{
		{Key: "userid", Value: req.UserId},
		{Key: "recordeddate", Value: bson.D{
			{Key: "$gte", Value: req.StartDate},
			{Key: "$lte", Value: req.EndDate},  // Use $lte instead of $lt
		}},
		{Key: "deletedat", Value: 0},
	}	

	fmt.Printf("Filter: %+v\n", filter)

	fmt.Println(req.StartDate, req.EndDate)
	cursor, err := m.Lifestyle.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var totalHeartRate float64
	var totalSleepDuration, totalTemperature float32
	var count int

	var averageLifestyle pb.WeeklySummary

	for cursor.Next(context.Background()) {
		var ls pb.WeeklySummary
		if err := cursor.Decode(&ls); err != nil {
			return nil, err
		}

		totalHeartRate += float64(ls.HeartRate)
		totalSleepDuration += ls.SleepDuration
		totalTemperature += ls.Temperature
		count++

		averageLifestyle.HeartRate = int32(totalHeartRate / float64(count))
		averageLifestyle.SleepDuration = float32(totalSleepDuration) / float32(count)
		averageLifestyle.Temperature = totalTemperature / float32(count)
		averageLifestyle.HeartRates = append(averageLifestyle.HeartRates, ls.HeartRate)
		averageLifestyle.SleepDurations = append(averageLifestyle.SleepDurations, ls.SleepDuration)
		averageLifestyle.Temperatures = append(averageLifestyle.Temperatures, ls.Temperature)
	}

	if count == 0 {
		return nil, fmt.Errorf("no data found for user %s within the given range", req.UserId)
	}

	fmt.Println(count)

	return &averageLifestyle, nil
}
