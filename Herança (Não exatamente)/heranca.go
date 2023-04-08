package main

import "fmt"

type Student struct {
	Person
	class  string
	passed bool
}

type Person struct {
	name    string
	surname string
	age     int
	heigth  float64
}

func main() {
	fmt.Println("------------ Herança ------------")

	person := Person{
		name:    "Vinicius",
		surname: "Vilela",
		age:     23,
		heigth:  1.70,
	}

	// Para inializar, é como se fosse uma composição, ou seja, passa-se uma struct como parametro
	student := Student{
		Person: person,
		class:  "Hello",
		passed: true,
	}
	fmt.Println("[student]:", student)

	// Porém, para acessar o objeto, não precisa referenciar 'person'
	fmt.Println("[student.name]", student.name)
}
