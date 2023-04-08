package main

import "fmt"

func main() {
	var (
		variavel1 string = "Hello"
		variavel2 string = "World"
	)
	fmt.Println("[TIPO DECLARADO]:", variavel1, variavel2)

	variavel3, variavel4 := "Hello", "World"
	fmt.Println("[TIPO INFERIDO]:", variavel3, variavel4)

	const constante1 string = "Hello World"
	fmt.Println("[CONSTANTE]:", constante1)

	variavel2, variavel1 = variavel1, variavel2
	fmt.Println("[SWAP]:", variavel1, variavel2)
}
