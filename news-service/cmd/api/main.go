package main

import (
	"github.com/neokofg/go-pet-detailed-microservices/news-service/internal/prelude"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/kafka"
)

func main() {
	logger := prelude.InitLogger()
	client := prelude.InitClient(logger)
	cleanup := kafka.InitConsumer(logger, client)
	prelude.InitServer(client, logger, cleanup)
}
