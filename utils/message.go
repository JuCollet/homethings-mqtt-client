package utils

import (
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Subscribe(topic string, handler func(message *kafka.Message)) error {
	config := make(kafka.ConfigMap)
	config["bootstrap.servers"] = os.Getenv("KAFKA_SERVER_HOST")
	config["group.id"] = "mqtt-client"

	c, err := kafka.NewConsumer(&config)

	if err != nil {
		return err
	}

	err = c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		return err
	}

	for {
		ev, err := c.ReadMessage(100 * time.Millisecond)

		if err != nil {
			continue
		}

		handler(ev)

	}

}
