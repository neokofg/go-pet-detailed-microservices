package main

import (
	"github.com/neokofg/go-pet-detailed-microservices/user-service/internal/app/events"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/internal/prelude"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/kafka"
)

func main() {
	logger := prelude.InitLogger()
	client := prelude.InitClient(logger)
	producer, cleanup := kafka.InitProducer(logger)
	events.UserDeleted(client, logger, producer)

	prelude.InitServer(client, logger, cleanup)
}
