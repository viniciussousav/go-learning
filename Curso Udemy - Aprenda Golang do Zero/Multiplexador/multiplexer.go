package main

import (
	"fmt"
	"time"
)

func produce() chan string {

	channel := make(chan string)

	go func() {
		for {
			channel <- "Hello World"
			time.Sleep(time.Second)
		}
	}()

	return channel
}

func consume(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-channel1:
				channel <- fmt.Sprintf("[CONSUMED] Message from channel1: %s", msg1)
			case msg2 := <-channel2:
				channel <- fmt.Sprintf("[CONSUMED] Message from channel2: %s", msg2)
			}
		}
	}()

	return channel
}

func main() {
	multiplexer := consume(produce(), produce())

	for value := range multiplexer {
		fmt.Println(value)
	}

}
