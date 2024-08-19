package app

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/mirjalilova/med-track/config"
	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	kafka "github.com/mirjalilova/med-track/pkg/kafka/consumer"
	"github.com/mirjalilova/med-track/pkg/websocket"
	st "github.com/mirjalilova/med-track/service/init"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	conn, err := st.ConnectDB()
	if err != nil {
		slog.Error("Failed to connect to the database: %v", err)
		return
	}

	medical := st.NewMedicalRecordService(conn)
	genetic := st.NewGeneticDataService(conn)
	lifestyle := st.NewLifeStyleService(conn)
	wearable := st.NewWearableDataService(conn)
	healthRecommendation := st.NewRecomendationService(conn)
	notification := st.NewNotificationService(conn)

	newServer := grpc.NewServer()
	pb.RegisterMedicalRecordServiceServer(newServer, medical)
	pb.RegisterGeneticDataServiceServer(newServer, genetic)
	pb.RegisterLifestyleServiceServer(newServer, lifestyle)
	pb.RegisterWearableDataServiceServer(newServer, wearable)
	pb.RegisterRecommendationServiceServer(newServer, healthRecommendation)
	pb.RegisterNotificationServiceServer(newServer, notification)

	// Kafka
	brokers := []string{"kafka_medtrack:9092"}
	cm := kafka.NewKafkaConsumerManager()
	RegisterConsumer(brokers, cm, medical)

	var wg sync.WaitGroup

	// Start gRPC server in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", cfg.MED_TRACK)
		if err != nil {
			slog.Error("Failed to listen on port %v: %v", cfg.MED_TRACK, err)
			return
		}
		slog.Info("Starting gRPC server on %v", cfg.MED_TRACK)
		if err := newServer.Serve(lis); err != nil {
			slog.Error("Failed to serve gRPC server: %v", err)
		}
	}()

	time.Sleep(2 * time.Second)

	// Start WebSocket server in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.HandleFunc("/ws", websocket.HandleWebSocket)
		slog.Info("WebSocket server starting on :7070")
		if err := http.ListenAndServe(":7070", nil); err != nil {
			slog.Error("ListenAndServe: %v", err)
		}
	}()

	wg.Wait()
}
