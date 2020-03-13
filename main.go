package main

import (
	c "github.com/davidogbiko/kafka-series/kafkaconsumer"
	p "github.com/davidogbiko/kafka-series/kafkaproducer"
)

func main() {
	// Producer example
	p.Producer()

	// Producer with callback example
	p.ProducerWithCallback()

	// Producer with keys
	p.ProducerWithKeys()

	// Consumer example
	c.Consumer()
}
