package kafkaproducer

import (
	"log"
	"strconv"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func ProducerWithCallback() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	topic := "first_topic"
	for i := 1; i <= 10; i++ {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte("hello world" + " " + strconv.Itoa(i)),
		}, nil)
	}

	// Golang has Event channels that can be polled for callbacks
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Fatal("error while producing", ev.TopicPartition.Error)
				} else {
					log.Print("received new metadata. \n",
						"Topic: ", *ev.TopicPartition.Topic, "\n",
						"Partition: ", ev.TopicPartition.Partition, "\n",
						"Offset: ", ev.TopicPartition.Offset, "\n",
					)
				}
			}
		}
	}()

	// flush waits for messages to be produced
	producer.Flush(15 * 100)
}
