package app

import (
	"net"

	pb "github.com/mirjalilova/med-track/genproto/health_sync"
	"github.com/mirjalilova/med-track/config"
	st "github.com/mirjalilova/med-track/service/init"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	conn, err := st.ConnectDB()
	if err != nil {
		slog.Error("")
		return
	}

	medical := st.NewMedicalRecordService(conn)
	genetic := st.NewGeneticDataService(conn)
	lifestyle := st.NewLifeStyleService(conn)
	wearable := st.NewWearableDataService(conn)
	healthRecommendation := st.NewRecomendationService(conn)

	newServer := grpc.NewServer()
	pb.RegisterMedicalRecordServiceServer(newServer, medical)
	pb.RegisterGeneticDataServiceServer(newServer, genetic)
	pb.RegisterLifestyleServiceServer(newServer, lifestyle)
	pb.RegisterWearableDataServiceServer(newServer, wearable)
	pb.RegisterRecommendationServiceServer(newServer, healthRecommendation)


	// Start gRPC server in a separate goroutine
	// go func() {
		lis, err := net.Listen("tcp", cfg.MED_TRACK)
		if err != nil {
			slog.Error("Failed to listen on port %v: %v", cfg.MED_TRACK, err)
		}
		slog.Info("Starting gRPC server on %v", cfg.MED_TRACK)
		if err := newServer.Serve(lis); err != nil {
			slog.Error("Failed to serve gRPC server: %v", err)
		}
	// }()
}
