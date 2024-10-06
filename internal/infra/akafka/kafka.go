package aKafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsume, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  servers,
		"group.id":           "user-group",
		"auto.offset.reset":  "earliest",
		"session.timeout.ms": 6000,
	})

	if err != nil {
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
