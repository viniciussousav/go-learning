package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func negative(x *int) *int {
	*x = -*x
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite um número para ter o sinal invertido: ")
	scanner.Scan()

	numberInput, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Número inválido: ", err)
	}

	result := negative(&numberInput)

	fmt.Println("Número digitado: ", numberInput)
	fmt.Println("Número negativo: ", *result)
}
