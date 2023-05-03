package main

import (
	"fmt"
	kafka_sample "go-kafka-example/kafka-sample"
)

func main() {

	producer := kafka_sample.Produce("hello-world")

	go func() {
		for events := range producer {
			fmt.Println(events)
		}
	}()

	consumer := kafka_sample.Consume("hello-world")

	for events := range consumer {
		fmt.Println(events)
	}
}
