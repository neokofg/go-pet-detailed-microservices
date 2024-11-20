package router

import (
	"github.com/gin-gonic/gin"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/grpc"
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/app/middleware"
	authProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/auth/v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type Router struct {
	app     *app.App
	cleanup []func()
}

func NewRouter(app *app.App) *Router {
	return &Router{
		app:     app,
		cleanup: make([]func(), 0),
	}
}

func (route *Router) InitRoutes(r *gin.Engine) {
	authClient, authCleanup := InitAuthService(route.app.Logger)
	route.cleanup = append(route.cleanup, authCleanup)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", route.app.Handlers.UserCommandHandler.Register)
			auth.POST("/login", route.app.Handlers.UserCommandHandler.Login)
			auth.POST("/logout", middleware.AuthMiddleware(authClient), route.app.Handlers.UserCommandHandler.Logout)
		}
		user := v1.Group("/user", middleware.AuthMiddleware(authClient))
		{
			user.GET("/me", route.app.Handlers.UserQueriesHandler.GetUser)
			user.PATCH("/me", route.app.Handlers.UserCommandHandler.UpdateUser)
		}
		news := v1.Group("/news")
		{
			news.GET("/feed")
			news.GET("/:id")
			news.POST("/new", middleware.AuthMiddleware(authClient))
			news.DELETE("/:id", middleware.AuthMiddleware(authClient))
			news.PATCH("/:id", middleware.AuthMiddleware(authClient))
		}
	}
}

func (route *Router) Close() {
	for _, cleanup := range route.cleanup {
		cleanup()
	}
}

func InitAuthService(logger *zap.Logger) (authProto.AuthServiceClient, func()) {
	authConn, err := grpc.InitGRPCClient(os.Getenv("AUTH_SERVICE_ADDR"))
	if err != nil {
		logger.Fatal("Failed to connect to catalog service", zap.Error(err))
	}

	cleanup := func() {
		if err := authConn.Close(); err != nil {
			logger.Error("Failed to close auth service connection", zap.Error(err))
		}
	}

	return authProto.NewAuthServiceClient(authConn), cleanup
}
