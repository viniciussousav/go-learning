package main

import "fmt"

type Address struct {
	street string
	number int64
}

type Person struct {
	name    string
	age     int
	heigth  float64
	address Address
}

func main() {
	fmt.Println("------- STRUCTS -------")

	// Criando com todos os parametros
	address1 := Address{"Rua Teste", 1}
	person1 := Person{"Vinicius", 23, 1.70, address1}
	fmt.Println("[person1]:", person1)

	// Parâmetros atribuídos após declaração
	var person2 Person
	person2.age = 23
	person2.name = "Vinicius"
	person2.heigth = 1.70
	fmt.Println("[person2]", person2)

	// Criando sem declarar valores
	var person3 Person
	fmt.Println("[person3]:", person3)

	// Criando declarando parametros parcialmente
	person4 := Person{name: "Vinicius"}
	fmt.Println("[person4]", person4)
}
