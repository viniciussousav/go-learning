package main

import (
	"fmt"
	"sync"
	"time"
)

func printTextLoop(text string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go printTextLoop("Hello World", &waitGroup)
	go printTextLoop("Hello Moon", &waitGroup)

	waitGroup.Wait()
}
