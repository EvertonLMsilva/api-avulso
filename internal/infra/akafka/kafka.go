package aKafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsume, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  servers,
		"group.id":           "user-group",
		"auto.offset.reset":  "earliest",
		"session.timeout.ms": 6000,
	})

	fmt.Printf("Error kafka", kafkaConsume)
	if err != nil {
		fmt.Printf("Error kafka", err)
		panic(err)
	}

	kafkaConsume.SubscriberTopics(topics, nil)
	for {
		msg, err := kafkaConsume.ReadMessage(-1)
		if err != nil {
			msgChan <- msg
		}
	}
}
