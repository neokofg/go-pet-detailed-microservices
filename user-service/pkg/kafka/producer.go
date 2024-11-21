package kafka

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"log"
	"os"
)

func InitProducer(logger *zap.Logger) (sarama.SyncProducer, func()) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_ADDR")}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	cleanup := func() {
		if err := producer.Close(); err != nil {
			logger.Error("Failed to close auth service connection", zap.Error(err))
		}
	}

	return producer, cleanup
}
