package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumAndSub(n1, n2 int64) (sum, sub int64) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite dois números")

	fmt.Print("Primeiro número: ")
	scanner.Scan()
	n1, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		panic("Número inválido")
	}

	fmt.Print("Segundo número: ")
	scanner.Scan()
	n2, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		panic("Número inválido")
	}

	sum, sub := sumAndSub(n1, n2)
	fmt.Printf("Soma: %d, Subtração: %d", sum, sub)
}
