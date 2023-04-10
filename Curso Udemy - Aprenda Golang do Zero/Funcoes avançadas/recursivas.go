package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonnaci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1 + fibonnaci(0)
	}

	return fibonnaci(n-2) + fibonnaci(n-1)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite um número para calcular a sequência de fibonnaci: ")
	scanner.Scan()

	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("O número informado não é válido")
	}

	for i := 0; i < number; i++ {
		fmt.Print(strconv.Itoa(fibonnaci(i)) + " ")
	}
}
