package handlers

import (
	authProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/auth/v1"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func InitHandlers(client *ent.Client, logger *zap.Logger) (*UserHandler, func()) {
	authConn, err := initGRPCClient(os.Getenv("AUTH_SERVICE_ADDR"))
	if err != nil {
		logger.Fatal("Failed to connect to catalog service", zap.Error(err))
	}

	cleanup := func() {
		if err := authConn.Close(); err != nil {
			logger.Error("Failed to close auth service connection", zap.Error(err))
		}
	}

	authClient := authProto.NewAuthServiceClient(authConn)

	userHandler := NewUserHandler(client, logger, authClient)

	return userHandler, cleanup
}

func initGRPCClient(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
