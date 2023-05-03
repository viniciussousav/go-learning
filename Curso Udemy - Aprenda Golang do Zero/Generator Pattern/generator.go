package main

import (
	"fmt"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Second * 2)
		}
	}()

	return channel
}

func main() {

	for i := range write("Hello world") {
		fmt.Println(i)
	}

}
