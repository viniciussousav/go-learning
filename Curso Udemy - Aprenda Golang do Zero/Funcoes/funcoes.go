package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sumAndMult(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	// Básico de função
	result := sum(1, 2)
	fmt.Println("[sum]", result)

	// Função declarada como variável
	var function = func(text string) string {
		fmt.Println("[function]", text)
		return text
	}
	fmt.Println("[main]", function("Hello World"))

	// Função com dois retornos
	resultSum, resultMul := sumAndMult(2, 6)
	fmt.Println("[resultSum]", resultSum)
	fmt.Println("[resultMult]", resultMul)

	// Ignorando o resultado do segundo retorno
	resultSum2, _ := sumAndMult(2, 3)
	fmt.Println("[resultSum2]", resultSum2)

	// Ignorando o resultado do primeiro retorno
	_, resultMult2 := sumAndMult(2, 3)
	fmt.Println("[resultMult2]", resultMult2)
}
