package kafkaproducer

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Producer is a base case kafka producer that publishes data to topics
func Producer() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	topic := "first_topic"
	for _, record := range []string{
		"hello world",
		"welcome to kafka series",
		"let's do something new"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(record),
		}, nil)

		// producer event channels can be polled for a more elegant check
		fmt.Printf("Record %q has been produced to topic: %s.\n", record, topic)
	}

	// flush waits for messages to be produced
	producer.Flush(15 * 100)
}
