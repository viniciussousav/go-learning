package auxiliar

import "fmt"

// Função utilizada para demonstrar que uma funcao de um mesmo pacote é visível para todos dentro dele
func escrever2(text string) {
	fmt.Printf("[escrever2]: %s", text)
}
