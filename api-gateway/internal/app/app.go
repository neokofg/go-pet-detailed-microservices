package app

import (
	"github.com/gin-gonic/gin"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/handlers"
	cache "github.com/neokofg/go-pet-detailed-microservices/api-gateway/pkg/redis"
	"go.uber.org/zap"
)

type App struct {
	router   *gin.Engine
	Logger   *zap.Logger
	Handlers *handlers.Handlers
	Cache    *cache.Cache
	cleanup  func()
}

func NewApp(logger *zap.Logger) *App {
	r := cache.NewCache()
	h, cleanup := handlers.InitHandlers(logger, r)
	return &App{
		Logger:   logger,
		Handlers: h,
		Cache:    r,
		cleanup:  cleanup,
	}
}

func (a *App) Close() {
	if a.cleanup != nil {
		a.cleanup()
	}
}
