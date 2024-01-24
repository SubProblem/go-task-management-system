package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/IBM/sarama"
)

type MessageConsumer struct {
	Consumer sarama.Consumer
}

type Message struct {
	TaskId int    `json:"taskId"`
	Task   string `json:"task"`
	UserId int    `json:"userId"`
}

func NewKafkaConsumer(groupId string) (*MessageConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}

	return &MessageConsumer{Consumer: consumer}, nil
}

func (c *MessageConsumer) Consume(wg *sync.WaitGroup) {
	defer wg.Done()

	partitionConsumer, err := c.Consumer.ConsumePartition("myTopic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	signals := make(chan os.Signal, 1)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var message Message
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				fmt.Println("Error decoding message: ", err)
				continue
			}
			fmt.Printf("Received message: TaskId=%d, Task=%s, UserId=%d\n",
				message.TaskId, message.Task, message.UserId)

		case err := <-partitionConsumer.Errors():
			fmt.Println("Error: ", err)

		case <-signals:
			break ConsumerLoop
		}
	}
}
