package apachkaf

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, chanMsg chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "fodase",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic("Fail while connecting with apache kafka")
	}
	kafkaConsumer.SubscribeTopics(topics, nil)
	for {
		message, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			chanMsg <- message
		}
	}
}
