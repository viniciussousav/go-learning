package auxiliar

import "fmt"

// Essa função escreve um texte e exibe no terminal utilizada em arquivo fora do pacote
func Escrever(text string) {
	fmt.Printf("[escrever]: %s", text)
	fmt.Println()
	escrever2(text)
}
