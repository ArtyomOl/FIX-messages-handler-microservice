package server

import (
	"FIX-messages-handler-API/fix"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
)

// Get messages from Kafka
func StartConsumer(client *redis.Client) {
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
		fmt.Println(string(message.Value))
		fix.AddFixMessage(client, string(message.Value))
	}
}

func SendToKafka(message string) error {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "fix-messages",
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
