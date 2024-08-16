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

type Notification struct {
	Notification *mongo.Collection
}

func NewNotification(mongoDb *mongo.Database) *Notification {
	return &Notification{
		Notification: mongoDb.Collection("notifications"),
	}
}

func (m *Notification) Create(req *pb.Notification) (*pb.Void, error) {
	notification := &pb.Notification{
		Id:        uuid.New().String(),
		Receiver:  req.Receiver,
		Message:   req.Message,
		CreatedAt: time.Now().Format(layout),
	}

	_, err := m.Notification.InsertOne(context.Background(), notification)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *Notification) Get(req *pb.GetById) (*pb.Notification, error) {
	var notification pb.Notification

	filter := bson.D{{Key: "id", Value: req.Id}}

	err := m.Notification.FindOne(context.TODO(), filter).Decode(&notification)
	if err == mongo.ErrNoDocuments {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (m *Notification) List(req *pb.GetAllNotificationReq) (*pb.GetAllNotificationRes, error) {
	res := &pb.GetAllNotificationRes{}

	filter := bson.D{{Key: "receiver", Value: req.Receiver}}

	findOptions := options.Find()
	if req.Filter != nil {
		findOptions.SetLimit(int64(req.Filter.Limit))
		findOptions.SetSkip(int64(req.Filter.Offset))
	}

	cur, err := m.Notification.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var notification pb.Notification
		if err := cur.Decode(&notification); err != nil {
			return nil, err
		}
		res.Notifications = append(res.Notifications, &notification)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	res.Count = int32(len(res.Notifications))
	return res, nil
}
