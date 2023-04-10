package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println("Digite quantos números quiser somar e digite 'calcular' para obter resultado")
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []float64

	for {
		scanner.Scan()

		input := scanner.Text()
		if input == "calcular" {
			break
		}

		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("O número digitado não é valido, insira novamente")
			continue
		}

		numbers = append(numbers, number)
	}

	result := sum(numbers...)
	fmt.Println("Resultado da soma:", result)
}
