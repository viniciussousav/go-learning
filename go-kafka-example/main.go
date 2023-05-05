package main

import (
	"fmt"
	kafka_sample "go-kafka-example/kafka-sample"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Set up a channel for handling Ctrl-C, etc
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	producer, err := kafka_sample.Produce("hello-world", &run)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	consumer, err := kafka_sample.Consume("hello-world", &run)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for {
		select {
		case <-signalChan:
			run = false
			go kafka_sample.ExitApplication()
		case p := <-producer:
			fmt.Println(p)
		case c := <-consumer:
			fmt.Println(c)
		default:
			continue
		}
	}
}
