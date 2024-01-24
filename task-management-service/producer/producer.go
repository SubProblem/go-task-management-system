package producer

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type MessageProducer struct {
	Producer sarama.SyncProducer
}

type Message struct {
	TaskId int    `json:"taskId"`
	Task   string `json:"task"`
	UserId int    `json:"userId"`
}

func NewKafkaProducer() (*MessageProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		return nil, err
	}

	return &MessageProducer{Producer: producer}, nil
}

func (p *MessageProducer) ProduceMessage(payload interface{}) error {

	topic := "myTopic"

	value, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := p.Producer.SendMessage(message)
	if err != nil {
		return err
	}

	fmt.Printf("Produced message to topic %s (partition %d) at offset %v\n", topic, partition, offset)

	return nil
}
