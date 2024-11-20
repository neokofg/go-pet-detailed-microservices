package handlers

import (
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent"
	"go.uber.org/zap"
)

func InitHandlers(client *ent.Client, logger *zap.Logger) (authHandler *AuthHandler) {
	authHandler = NewAuthHandler(client, logger)

	return authHandler
}
