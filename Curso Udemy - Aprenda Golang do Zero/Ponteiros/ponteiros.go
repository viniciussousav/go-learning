package main

import "fmt"

func main() {
	fmt.Println("--------- PONTEIROS ---------")

	var number1 int = 1
	var number2 int = number1
	fmt.Println("[number1 number2]", number1, number2)

	number1 = 20
	fmt.Println("'number1' foi atualizada para 20")

	// A atualização não reflete no number2, pois o valor foi passado por cópia
	fmt.Println("[number1 number2]", number1, number2)

	var number3 int = 3
	var number4 *int = &number3
	fmt.Println("[number3 number4]", number3, *number4)

	number3 = 30
	fmt.Println("'number3' foi atualizada para 30")

	// A atualização refletiu, pois number4 aponta para o valor na memória de number3
	fmt.Println("[number3 number4]", number3, *number4)
}
