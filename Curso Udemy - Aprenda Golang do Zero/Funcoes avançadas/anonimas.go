package main

import (
	"fmt"
)

func main() {
	func(name string) {
		fmt.Println("Hello", name)
	}("Vinicius")

	received := func(name string) string {
		return fmt.Sprintf("Recebido -> %s", name)
	}("Vinicius")
	fmt.Println(received)
}
