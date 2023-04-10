package main

import (
	"errors"
	"fmt"
)

func main() {

	var (
		numberIntInfered int   = 0
		numberInt8       int8  = 8  // inteiro de até 8 bits
		numberInt16      int16 = 16 // inteiro de até 16 bits
		numberInt32      int32 = 32 // inteiro de até 32 bits
		numberInt64      int64 = 64 // inteiro de até 64 bits

		/*
			Os exemplos abaixo não compilam:

			invalidNumberInt8  int8  = 1000000000000000
			invalidNumberInt16 int16 = 1000000000000000
			invalidNumberInt32 int32 = 1000000000000000
			invalidNumberInt64 int64 = 1000000000000000
		*/
	)

	fmt.Println("[Inteiros]", numberIntInfered, numberInt8, numberInt16, numberInt32, numberInt64)

	var (
		numberFloat32 float32 = 8  // número real com até 32 bits
		numberFloat64 float64 = 12 // número real com até 64 bits
	)

	fmt.Println("[Reais]", numberFloat32, numberFloat64)

	var stringExample1 string = "Texto 1" // string declarada de forma explícita
	fmt.Println("[String 1]", stringExample1)

	stringExample2 := "Texto 2" // string com inferência de tipo
	fmt.Println("[String 2]", stringExample2)

	char := 'B' // o tipo 'char' não existe em GO, ao printar essa variável é exibido o valor do caractere na tabela ASCII
	fmt.Println("[Caractere]", char)

	var stringDefaultValue string // valor padrão de string é uma string vazia
	fmt.Println("[String Valor Padrão]", stringDefaultValue)

	var intDefaultValue int // valor padrão de qualquer tipo numérico é 0
	fmt.Println("[Int Valor Padrão]", intDefaultValue)

	var boolDefaultValue bool // valor padrão de tipo boleano é false
	fmt.Println("[Bool Valor Padrão]", boolDefaultValue)

	var errorDefaultValue error // valor padrão de tipo 'error' é 'nil'
	fmt.Println("[Error Valor Padrão]", errorDefaultValue)

	var errorExample error = errors.New("Error interno")
	fmt.Println("[Error Exemplo]", errorExample)
}
