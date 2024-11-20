package servers

import (
	"fmt"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/internal/app/handlers"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
)

func InitGrpcServer(client *ent.Client, logger *zap.Logger) (*grpc.Server, func()) {
	userHandler, cleanup := handlers.InitHandlers(client, logger)

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userHandler)

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
