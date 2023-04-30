package main

import (
	"fmt"
	"time"
)

func ReceiveMessage(id int, channel *chan string) {
	for value := range *channel {
		fmt.Printf("[%d - GOROUTINE] RECEIVED: \"%s\" \n", id, value)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan string, 4)

	for i := 0; i < 4; i++ {
		fmt.Printf("[%d - GOROUTINE] STARTED \n", i)
		go ReceiveMessage(i, &channel)
		time.Sleep(time.Second)
	}

	fmt.Println("CONSUMING WILL START IN 5 SECONDS...")
	time.Sleep(time.Second * 5)

	for i := 0; i < 20; i++ {
		channel <- fmt.Sprintf("Message %d", i)
	}

	fmt.Println("FINISHING PROGRAM IN 5 SECONDS...")
	time.Sleep(time.Second * 5)
	fmt.Println("PROGRAM FINISHED")
}
