package prelude

import (
	"context"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/internal/prelude/servers"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer(client *ent.Client, logger *zap.Logger) {
	defer logger.Info("Servers exited properly")

	grpcServer := servers.InitGrpcServer(client, logger)
	defer grpcServer.GracefulStop()

	ginServer := servers.InitGinServer(logger)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := ginServer.Shutdown(ctx); err != nil {
			logger.Fatal("HTTP server forced to shutdown", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down servers...")
}
