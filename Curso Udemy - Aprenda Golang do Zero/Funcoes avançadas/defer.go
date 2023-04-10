package main

import (
	"fmt"
	"time"
)

func firstFunction(iterations int) {
	for i := int32(0); i < int32(iterations); i++ {
		fmt.Println("Executando função 1")
		time.Sleep(time.Second)
	}
}

func secondFunction(iterations int) {
	for i := int32(0); i < int32(iterations); i++ {
		fmt.Println("Executando função 2")
		time.Sleep(time.Second)
	}
}

func main() {
	defer firstFunction(10)
	secondFunction(10)
}
