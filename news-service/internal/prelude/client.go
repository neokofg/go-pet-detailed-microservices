package prelude

import (
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent"
	"go.uber.org/zap"
	"os"
)

func InitClient(logger *zap.Logger) *ent.Client {
	client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
		panic(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("Failed to create schema", zap.Error(err))
		panic(err)
	}

	return client
}
