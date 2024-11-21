package prelude

import (
	"context"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/internal/prelude/servers"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer(client *ent.Client, logger *zap.Logger, closeProducer func()) {
	defer logger.Info("Servers exited properly")
	defer closeProducer()
	defer client.Close()
	grpcServer, cleanup := servers.InitGrpcServer(client, logger)
	defer grpcServer.GracefulStop()
	defer cleanup()

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
