package main

import (
	"fmt"
	"time"
)

func produceToFirstchannel(channel chan string) {
	for {
		time.Sleep(time.Second * 1)
		channel <- "Canal 1"
	}
}

func produceToSecondChannel(channel chan string) {
	for {
		time.Sleep(time.Second * 4)
		channel <- "Canal 2"
	}
}

func main() {

	firstChannel, secondChannel := make(chan string), make(chan string)

	go produceToFirstchannel(firstChannel)
	go produceToSecondChannel(secondChannel)

	for {
		select {
		case firstChannelStr := <-firstChannel:
			fmt.Println(firstChannelStr)
		case secondChannelStr := <-secondChannel:
			fmt.Println(secondChannelStr)
		}
	}
}
