package events

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"go.uber.org/zap"
)

func UserDeleted(client *ent.Client, logger *zap.Logger, producer sarama.SyncProducer) {
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(ent.OpDelete) {
				id, _ := m.Field("id")

				payload := struct {
					UserID uuid.UUID `json:"user_id"`
				}{
					UserID: id.(uuid.UUID),
				}

				jsonData, err := json.Marshal(payload)
				if err != nil {
					logger.Error("Failed to marshal message", zap.Error(err))
					return next.Mutate(ctx, m)
				}

				msg := &sarama.ProducerMessage{
					Topic: "user.deleted.news",
					Value: sarama.StringEncoder(jsonData),
				}

				if _, _, err := producer.SendMessage(msg); err != nil {
					logger.Error("Can't produce message", zap.Error(err))
				}

				msg = &sarama.ProducerMessage{
					Topic: "user.deleted.auth",
					Value: sarama.StringEncoder(jsonData),
				}

				if _, _, err := producer.SendMessage(msg); err != nil {
					logger.Error("Can't produce message", zap.Error(err))
				}
			}

			return next.Mutate(ctx, m)
		})
	})
}
