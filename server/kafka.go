package server

import (
	"FIX-messages-handler-API/fix"

	"github.com/IBM/sarama"
)

// Get messages from Kafka
func StartConsumer() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("fix-messages", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		fix.AddFixMessage(string(message.Value))
	}
}
