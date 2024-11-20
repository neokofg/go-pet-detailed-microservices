package handlers

import (
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/grpc"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/handlers/commands"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/handlers/queries"
	userProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"go.uber.org/zap"
	"os"
)

type Handlers struct {
	UserCommandHandler commands.UserCommandsHandler
	UserQueriesHandler queries.UserQueriesHandler
}

func InitHandlers(logger *zap.Logger) (*Handlers, func()) {
	userConn, err := grpc.InitGRPCClient(os.Getenv("USER_SERVICE_ADDR"))
	if err != nil {
		logger.Fatal("Failed to connect to user service", zap.Error(err))
		panic("Failed to connect to user service")
	}

	cleanup := func() {
		if err := userConn.Close(); err != nil {
			logger.Error("Failed to close user service connection", zap.Error(err))
		}
	}

	userClient := userProto.NewUserServiceClient(userConn)

	userCommandsHandler := *commands.NewUserCommandsHandler(logger, userClient)
	userQueriesHandler := *queries.NewUserQueriesHandler(logger, userClient)

	return &Handlers{
		UserCommandHandler: userCommandsHandler,
		UserQueriesHandler: userQueriesHandler,
	}, cleanup
}
