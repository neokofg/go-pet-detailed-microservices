package main

import (
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/internal/prelude"
)

func main() {
	logger := prelude.InitLogger()
	client := prelude.InitClient(logger)

	prelude.InitServer(client, logger)
}
