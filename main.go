package main

import "github.com/davidogbiko/kafka-series/kafkaproducer"

func main() {
	// Producer example
	kafkaproducer.Producer()

	// Producer with callback example
	kafkaproducer.ProducerWithCallback()

	// Producer with keys
	kafkaproducer.ProducerWithKeys()
}
