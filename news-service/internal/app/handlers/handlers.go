package handlers

import (
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/ent"
	userProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func InitHandlers(client *ent.Client, logger *zap.Logger) (*NewsHandler, func()) {
	userConn, err := initGRPCClient(os.Getenv("USER_SERVICE_ADDR"))
	if err != nil {
		logger.Fatal("Failed to connect to catalog service", zap.Error(err))
	}

	cleanup := func() {
		if err := userConn.Close(); err != nil {
			logger.Error("Failed to close auth service connection", zap.Error(err))
		}
	}

	userClient := userProto.NewUserServiceClient(userConn)

	newsHandler := NewNewsHandler(userClient, client, logger)

	return newsHandler, cleanup
}

func initGRPCClient(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
