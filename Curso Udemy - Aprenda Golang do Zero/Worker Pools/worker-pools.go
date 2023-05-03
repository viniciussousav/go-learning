package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return fib(0) + 1
	default:
		return fib(n-2) + fib(n-1)
	}
}

func worker(numbers <-chan int, results chan<- int) {
	for number := range numbers {
		results <- fib(number)
	}
}

func main() {

	n := 43

	numbers := make(chan int, n)
	results := make(chan int, n)

	startTime := time.Now()

	go worker(numbers, results)

	for i := 0; i < n; i++ {
		numbers <- i
	}
	close(numbers)

	for i := 0; i < n; i++ {
		result := <-results
		fmt.Printf("fib(%d) = %d \n", i, result)
	}

	endTime := time.Now()
	totalTime := endTime.Sub(startTime)
	fmt.Printf("TOTAL TIME: %.2f \n", totalTime.Seconds())

}
