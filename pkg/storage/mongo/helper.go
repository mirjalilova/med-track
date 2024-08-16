package mongo

import (
	"context"
	"fmt"
	"time"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"go.mongodb.org/mongo-driver/bson"
)

// CalculateStressLevel calculates the stress level based on various health metrics.
func CalculateStressLevel(heartRate int32, sleepDuration float64, stepCount int32) int {
	stressLevel := 0

	if heartRate > 80 {
		stressLevel += 3
	} else if heartRate > 70 {
		stressLevel += 2
	} else {
		stressLevel += 1
	}

	if sleepDuration < 6 {
		stressLevel += 3
	} else if sleepDuration < 7 {
		stressLevel += 2
	} else {
		stressLevel += 1
	}

	if stepCount < 5000 {
		stressLevel += 3
	} else if stepCount < 10000 {
		stressLevel += 2
	} else {
		stressLevel += 1
	}

	return stressLevel
}

func (m *WearableData) CalculateDailyAverages(userId string, date string) (*pb.WearableData, error) {
	parsedDate, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return nil, err
	}
	startOfDay := parsedDate.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	filter := bson.D{
		{Key: "userid", Value: userId},
		{Key: "recordedtimestamp", Value: bson.D{
			{Key: "$gte", Value: startOfDay.Format("2006-01-02 15:04:05")},
			{Key: "$lt", Value: endOfDay.Format("2006-01-02 15:04:05")},
		}},
		{Key: "deletedat", Value: 0},
	}

	cursor, err := m.WearableData.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var totalHeartRate float64
	var totalSleepDuration, totalCaloriesBurned, totalTemperature float32
	var count int

	for cursor.Next(context.Background()) {
		var wd pb.WearableData
		if err := cursor.Decode(&wd); err != nil {
			return nil, err
		}

		totalHeartRate += float64(wd.HeartRate)
		totalSleepDuration += wd.SleepDuration
		totalCaloriesBurned += wd.CaloriesBurned
		totalTemperature += wd.Temperature
		count++
	}

	if count == 0 {
		return nil, fmt.Errorf("no data found for user %s on date %s", userId, date)
	}

	averageData := &pb.WearableData{
		HeartRate:      int32(totalHeartRate / float64(count)),
		SleepDuration:  float32(totalSleepDuration) / float32(count),
		CaloriesBurned: float32(totalCaloriesBurned) / float32(count),
		Temperature:    totalTemperature / float32(count),
	}

	return averageData, nil
}

func (m *WearableData) CalculateWeeklySummary(userId string, startDate string) (*pb.WeeklySummary, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}

	end := start.AddDate(0, 0, 7)
	filter := bson.D{
		{Key: "userid", Value: userId},
		{Key: "recordedtimestamp", Value: bson.D{
			{Key: "$gte", Value: start},
			{Key: "$lt", Value: end},
		}},
		{Key: "deletedat", Value: 0},
	}

	cursor, err := m.WearableData.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var totalHeartRate float64
	var totalSleepDuration, totalCaloriesBurned, totalTemperature float32
	var heartRates []int32
	var sleepDurations, temperatures []float32
	var count int

	for cursor.Next(context.Background()) {
		var wd pb.WearableData
		if err := cursor.Decode(&wd); err != nil {
			return nil, err
		}

		totalHeartRate += float64(wd.HeartRate)
		heartRates = append(heartRates, wd.HeartRate)

		totalSleepDuration += wd.SleepDuration
		sleepDurations = append(sleepDurations, float32(wd.SleepDuration))

		totalCaloriesBurned += wd.CaloriesBurned
		totalTemperature += wd.Temperature
		temperatures = append(temperatures, float32(wd.Temperature))
		count++
	}

	if count == 0 {
		return nil, fmt.Errorf("no data found for user %s in the week starting from %s", userId, startDate)
	}

	weeklySummary := &pb.WeeklySummary{
		HeartRate:      int32(totalHeartRate / float64(count)),
		HeartRates:     heartRates,
		SleepDuration:  float32(totalSleepDuration) / float32(count),
		SleepDurations: sleepDurations,
		Temperature:    totalTemperature / float32(count),
		Temperatures:   temperatures,
	}

	return weeklySummary, nil
}
