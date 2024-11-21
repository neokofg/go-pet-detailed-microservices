package main

import (
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/internal/prelude"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/kafka"
)

func main() {
	logger := prelude.InitLogger()
	client := prelude.InitClient(logger)
	cleanup := kafka.InitConsumer(logger, client)

	prelude.InitServer(client, logger, cleanup)
}
