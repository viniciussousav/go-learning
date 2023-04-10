package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getHeight(scanner bufio.Scanner) float64 {
	fmt.Println("Insira a sua altura (ex: 1.75)")
	fmt.Print("-> ")
	scanner.Scan()
	heightInputString := scanner.Text()
	height, err := strconv.ParseFloat(heightInputString, 64)
	if err != nil {
		panic("Altura inválida: " + err.Error())
	}
	return height
}

func getWeight(scanner bufio.Scanner) float64 {
	fmt.Println("Insira o seu peso")
	fmt.Print("-> ")
	scanner.Scan()
	weightInputString := scanner.Text()

	weigth, err := strconv.ParseFloat(weightInputString, 64)
	if err != nil {
		panic("Peso inválido: " + err.Error())
	}
	return weigth
}

func printImcResult(weigth float64, height float64) {
	if imc := weigth / math.Pow(height, 2); imc < 18.5 {
		fmt.Printf("IMC = %f, Você está abaixo do peso", imc)
	} else if imc > 24.9 {
		fmt.Printf("IMC = %f, Você está acima do peso", imc)
	} else {
		fmt.Printf("IMC = %f, Você está no peso ideal", imc)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo, vamos verificar seu IMC")

	height := getHeight(*scanner)
	weigth := getWeight(*scanner)

	printImcResult(weigth, height)
}
