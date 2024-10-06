package aKafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsume, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "user-group",
		"auto.offset.reset": "earliest",
		"acks":              "all",
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
