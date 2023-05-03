package kafka_sample

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Produce(topic string) (result <-chan string) {

	channel := make(chan string)

	conf := ReadConfig("./getting_started.properties")

	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					channel <- fmt.Sprintf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					channel <- fmt.Sprintf("Produced event to topic %s: value = %s\n", *ev.TopicPartition.Topic, string(ev.Value))
				}
			}
		}
	}()

	go func() {
		for {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            nil,
				Value:          []byte("Hello World!"),
			}, nil)

			time.Sleep(time.Second * 2)
		}
	}()

	return channel
}
