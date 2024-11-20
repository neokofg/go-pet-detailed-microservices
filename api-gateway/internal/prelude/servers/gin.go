package servers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/middleware"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/router"
	"go.uber.org/zap"
	"net/http"
)

func InitGinServer(logger *zap.Logger) (*http.Server, func()) {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.CORSMiddleware(),
		middleware.PrometheusMiddleware(),
		middleware.RequestLoggerMiddleware(logger),
	)

	mainApp := app.NewApp(logger)

	mainRouter := router.NewRouter(mainApp)
	mainRouter.InitRoutes(r)

	cleanup := func() {
		mainApp.Close()
		mainRouter.Close()
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		logger.Info("Starting server...")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	return httpServer, cleanup
}
