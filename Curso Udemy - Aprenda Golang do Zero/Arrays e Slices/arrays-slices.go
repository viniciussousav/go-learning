package main

import "fmt"

func main() {
	fmt.Println("-------------- ARRAYS E SLICE1S --------------")

	// obrigatoriamente especificar o tamanho do array
	var array1 [5]int
	array1[0] = 1
	fmt.Println("[array1]", array1)

	// inicializando durante declaração
	array2 := [5]string{"Hello", "World", "Is", "The", "Law"}
	fmt.Println("[array2]", array2)

	// tamanho de acordo com a inicialização
	array3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("[array3]", array3)

	// Declaração de um slice
	slice1 := []int{2, 4, 6, 8, 10}
	fmt.Println("[slice1]", slice1)

	// adicionando elementos ao slice
	slice1 = append(slice1, 12, 14)
	fmt.Println("[slice1]", slice1)

	// criando slice a partir de array já existente
	slice2 := array3[1:5]
	fmt.Println("[slice2]", slice2)

	// o slice2 reflete as mudanças em array3 pois acontece uma referência
	array3[1] = 2
	fmt.Println("[slice2]", slice2)

	// array internos
	slice3 := make([]int, 10, 15) // inicializa um slice, especificando o tipo do slice, o tamanho inicial e a capacidade
	fmt.Println("[slice3]", slice3)
	fmt.Println("[slice3 len]", len(slice3)) // len retorna a quantidade de itens presente no slice
	fmt.Println("[slice3 cap]", cap(slice3)) // cap retorna a capacidade máxima de itens de um slice

	slice3 = append(slice3, 1, 2, 3, 4, 5, 6) // inserindo quantidade para ultrapassar a capacidade atual
	fmt.Println("[slice3]", slice3)
	fmt.Println("[slice3 len]", len(slice3)) // len retorna a quantidade de itens presente no slice
	fmt.Println("[slice3 cap]", cap(slice3)) // toda vez que a capacidade do slice não for suficiente, ela será dobrada
}
