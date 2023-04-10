package main

import "fmt"

func closure() func() {

	message := "Função closure"

	function := func() { fmt.Println(message) }
	return function
}

func main() {

	message := "Função main"
	fmt.Println(message)

	newFunction := closure()
	newFunction()

}
