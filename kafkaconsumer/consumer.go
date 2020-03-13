package kafkaconsumer

import (
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Consumer reads messages from a subsribed topic.
func Consumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "application_group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{"first_topic"}, nil)

	for {
		record, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, record)
		} else {
			log.Printf("Message on %s: %s\n", record.TopicPartition, string(record.Value))
		}
	}

	consumer.Close()
}
