package main

import "fmt"

func handleInvalidArg() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado do erro: ", r)
	}
}

func div(x int, y int) int {
	defer handleInvalidArg()

	if y == 0 {
		panic("Zero não é um divisor válido")
	}

	return x / y
}

func main() {
	fmt.Println(div(10, 0))
}
