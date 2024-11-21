package servers

import (
	"fmt"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/internal/app/handlers"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/ent"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
)

func InitGrpcServer(client *ent.Client, logger *zap.Logger) (*grpc.Server, func()) {
	newsHandler, cleanup := handlers.InitHandlers(client, logger)

	grpcServer := grpc.NewServer()
	proto.RegisterNewsServiceServer(grpcServer, newsHandler)

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	grpcAddr := fmt.Sprintf(":%s", os.Getenv("GRPC_PORT"))
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	go func() {
		logger.Info("Starting gRPC server", zap.String("addr", grpcAddr))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve gRPC", zap.Error(err))
		}
	}()

	return grpcServer, cleanup
}
