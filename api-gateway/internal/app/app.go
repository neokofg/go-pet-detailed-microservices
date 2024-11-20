package app

import (
	"github.com/gin-gonic/gin"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/handlers"
	"go.uber.org/zap"
)

type App struct {
	router   *gin.Engine
	Logger   *zap.Logger
	Handlers *handlers.Handlers
	cleanup  func()
}

func NewApp(logger *zap.Logger) *App {
	handlers, cleanup := handlers.InitHandlers(logger)
	return &App{
		Logger:   logger,
		Handlers: handlers,
		cleanup:  cleanup,
	}
}

func (a *App) Close() {
	if a.cleanup != nil {
		a.cleanup()
	}
}
