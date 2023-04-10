package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     uint
	hasCar  bool
}

func (p Person) fullName() string {
	return p.name + " " + p.surname
}

func (p Person) isAdult() bool {
	return p.age >= 18
}

func (p *Person) setHasCar(value bool) {
	p.hasCar = value
}

func main() {
	person := Person{"Vinicius", "Vilela", 23, false}
	fmt.Printf("%s é maior de idade? %t\n", person.fullName(), person.isAdult())
	fmt.Printf("%s tem carro? %t\n", person.fullName(), person.hasCar)

	person.setHasCar(true)

	fmt.Printf("%s comprou um carro? %t", person.fullName(), person.hasCar)
}
