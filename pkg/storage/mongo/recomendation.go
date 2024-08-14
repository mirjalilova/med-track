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

type Recommendation struct {
	Recommendation *mongo.Collection
}

func NewRecommendation(mongoDb *mongo.Database) *Recommendation {
	return &Recommendation{
		Recommendation: mongoDb.Collection("recommendations"),
	}
}

func (m *Recommendation) Create(req *pb.RecommendationCreate) (*pb.Void, error) {
	recommendation := &pb.Recommendation{
		Id:                 uuid.New().String(),
		UserId:             req.UserId,
		RecommendationType: req.RecommendationType,
		Description:        req.Description,
		Priority:           req.Priority,
		RecommendatedBy:    req.RecommendatedBy,
		CreatedAt:          time.Now().Format(layout),
		UpdatedAt:          time.Now().Format(layout),
		DeletedAt:          0,
	}

	_, err := m.Recommendation.InsertOne(context.Background(), recommendation)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *Recommendation) Update(req *pb.RecommendationUpdate) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	updateFields := bson.D{
		{Key: "updatedat", Value: time.Now().Format(layout)},
	}

	if req.RecommendationType != "" {
		updateFields = append(updateFields, bson.E{Key: "recommendationtype", Value: req.RecommendationType})
	}
	if req.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: req.Description})
	}
	if req.Priority!= 0 {
        updateFields = append(updateFields, bson.E{Key: "priority", Value: req.Priority})
    }

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err := m.Recommendation.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *Recommendation) Delete(req *pb.GetById) (*pb.Void, error) {
	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deletedat", Value: time.Now().Unix()},
		}},
	}

	_, err := m.Recommendation.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (m *Recommendation) Get(req *pb.GetById) (*pb.RecommendationRes, error) {
	res := &pb.RecommendationRes{}

	filter := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "deletedat", Value: 0},
	}

	err := m.Recommendation.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *Recommendation) List(req *pb.GetAllRecommendationReq) (*pb.GetAllRecommendation, error) {
	res := &pb.GetAllRecommendation{}

	filter := bson.D{{Key: "deletedat", Value: 0}}
	if req.UserId != "" {
		filter = append(filter, bson.E{Key: "userid", Value: req.UserId})
	}
	if req.RecommendatedBy != "" {
		filter = append(filter, bson.E{Key: "recommendatedby", Value: req.RecommendatedBy})
	}
	if req.Priority!= 0 {
        filter = append(filter, bson.E{Key: "priority", Value: req.Priority})
    }

	findOptions := options.Find()
	if req.Filter != nil {
		findOptions.SetLimit(int64(req.Filter.Limit))
		findOptions.SetSkip(int64(req.Filter.Offset))
	}

	cur, err := m.Recommendation.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var doc pb.RecommendationRes
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		res.Recommendations = append(res.Recommendations, &doc)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
