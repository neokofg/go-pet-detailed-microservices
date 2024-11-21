package kafka

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent/personalaccesstoken"
	"go.uber.org/zap"
	"log"
	"os"
)

type UserDeletedMessage struct {
	UserID uuid.UUID `json:"user_id"`
}

func InitConsumer(logger *zap.Logger, client *ent.Client) func() {
	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_ADDR")}, nil)
	if err != nil {
		log.Fatal(err)
	}

	partitionConsumer, err := consumer.ConsumePartition("user.deleted.auth", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for message := range partitionConsumer.Messages() {
			var msg UserDeletedMessage
			if err := json.Unmarshal(message.Value, &msg); err != nil {
				logger.Error("Failed to unmarshal message", zap.Error(err))
				continue
			}
			if _, err := client.PersonalAccessToken.Delete().
				Where(personalaccesstoken.UserIDEQ(msg.UserID)).
				Exec(context.Background()); err != nil {
				logger.Error("Failed to delete pat", zap.Error(err))
			}
		}
	}()

	cleanup := func() {
		if err := partitionConsumer.Close(); err != nil {
			logger.Error("Failed to close partition connection", zap.Error(err))
		}
		if err := consumer.Close(); err != nil {
			logger.Error("Failed to close kafka connection", zap.Error(err))
		}
	}

	return cleanup
}
