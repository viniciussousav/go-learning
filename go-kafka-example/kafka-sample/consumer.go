package kafka_sample

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consume(topic string) chan string {

	channel := make(chan string)

	conf := ReadConfig("./getting_started.properties")

	conf["group.id"] = "go-kafka-example"
	conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	err = c.SubscribeTopics([]string{topic}, nil)

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true

	go func() {
		for run {
			select {
			case sig := <-sigchan:
				channel <- fmt.Sprintf("Caught signal %v: terminating\n", sig)
				run = false
				close(channel)
			default:
				ev, err := c.ReadMessage(100 * time.Millisecond)
				if err != nil {
					// Errors are informational and automatically handled by the consumer
					continue
				}
				channel <- fmt.Sprintf("Consumed event from topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			}
		}
	}()

	return channel
}
