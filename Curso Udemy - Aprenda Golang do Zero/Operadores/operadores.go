package main

import "fmt"

func main() {

	// Operadores aritméticos
	fmt.Println("\n------- OPERADORES ARITMÉTICOS -------")

	sum := 1 + 2
	sub := 2 - 1
	div := 4 / 2
	mult := 2 * 6
	mod := 7 % 2

	fmt.Println("[sum]", sum)
	fmt.Println("[sub]", sub)
	fmt.Println("[div]", div)
	fmt.Println("[mult]", mult)
	fmt.Println("[mod]", mod)

	// Operadores relacionais
	fmt.Println("\n------- OPERADORES RELACIONAIS -------")

	fmt.Println("[1 > 2]", 1 > 2)
	fmt.Println("[1 >= 2]", 1 >= 2)
	fmt.Println("[1 == 2]", 1 == 2)
	fmt.Println("[1 < 2]", 1 < 2)
	fmt.Println("[1 <= 2]", 1 <= 2)
	fmt.Println("[1 != 2]", 1 != 2)

	// Operadores lógicos
	fmt.Println("\n------- OPERADORES LÓGICOS -------")

	trueValue, falseValue := true, false

	fmt.Println("[true && false]", trueValue && falseValue)
	fmt.Println("[true || false]", trueValue || falseValue)
	fmt.Println("[!true]", !trueValue)

}
