package prelude

import (
	"context"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/prelude/servers"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer(logger *zap.Logger) {
	defer logger.Info("Servers exited properly")

	ginServer, cleanup := servers.InitGinServer(logger)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cleanup()

		if err := ginServer.Shutdown(ctx); err != nil {
			logger.Fatal("HTTP server forced to shutdown", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down servers...")
}
