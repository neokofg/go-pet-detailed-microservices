package handlers

import (
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/grpc"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/handlers/commands"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/handlers/queries"
	cache "github.com/neokofg/go-pet-detailed-microservices/api-gateway/pkg/redis"
	newsProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	userProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"go.uber.org/zap"
	"os"
)

type Handlers struct {
	UserCommandHandler commands.UserCommandsHandler
	UserQueriesHandler queries.UserQueriesHandler
	NewsCommandHandler commands.NewsCommandsHandler
	NewsQueriesHandler queries.NewsQueriesHandler
}

func InitHandlers(logger *zap.Logger, cache *cache.Cache) (*Handlers, func()) {
	userConn, err := grpc.InitGRPCClient(os.Getenv("USER_SERVICE_ADDR"))
	if err != nil {
		logger.Fatal("Failed to connect to user service", zap.Error(err))
		panic("Failed to connect to user service")
	}

	userClient := userProto.NewUserServiceClient(userConn)

	userCommandsHandler := *commands.NewUserCommandsHandler(logger, userClient)
	userQueriesHandler := *queries.NewUserQueriesHandler(logger, userClient)

	newsConn, err := grpc.InitGRPCClient(os.Getenv("NEWS_SERVICE_ADDR"))
	if err != nil {
		logger.Fatal("Failed to connect to news service", zap.Error(err))
		panic("Failed to connect to user service")
	}

	newsClient := newsProto.NewNewsServiceClient(newsConn)

	newsCommandsHandler := *commands.NewNewsCommandsHandler(logger, newsClient)
	newsQueriesHandler := *queries.NewNewsQueriesHandler(logger, cache, newsClient)

	cleanup := func() {
		if err := newsConn.Close(); err != nil {
			logger.Error("Failed to close user service connection", zap.Error(err))
		}
		if err := userConn.Close(); err != nil {
			logger.Error("Failed to close user service connection", zap.Error(err))
		}
	}

	return &Handlers{
		UserCommandHandler: userCommandsHandler,
		UserQueriesHandler: userQueriesHandler,
		NewsCommandHandler: newsCommandsHandler,
		NewsQueriesHandler: newsQueriesHandler,
	}, cleanup
}
