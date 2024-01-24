package main

import (
	"log"
	"subproblem/notification-service/consumer"
	"sync"
)

func main() {

	topic := "myTopic"
	kafkaConsumer, err := consumer.NewKafkaConsumer(topic)

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	kafkaConsumer.Consume(&wg)
}