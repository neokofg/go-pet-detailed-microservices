package main

import (
	"github.com/neokofg/go-pet-detailed-microservices/api-gateway/internal/prelude"
)

func main() {
	logger := prelude.InitLogger()
	prelude.InitServer(logger)
}
