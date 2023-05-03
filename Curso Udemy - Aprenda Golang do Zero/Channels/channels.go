package main

import (
	"fmt"
	"time"
)

func sendToChannel(value string, channel chan string) {
	for i := 0; i < 3; i++ {
		channel <- value
		time.Sleep(time.Second)
	}
	close(channel)
}

func receiveFromChannel(channel chan string) {
	for value := range channel {
		fmt.Println(value)
	}
}

func main() {
	channel := make(chan string)
	go sendToChannel("Hello World", channel)
	receiveFromChannel(channel)
}
