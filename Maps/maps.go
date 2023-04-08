package main

import "fmt"

func main() {
	fmt.Println("-------------------- MAPS --------------------")

	users := map[int]string{
		1: "Vinicius",
		2: "João",
		3: "Maria",
		4: "José",
	}
	fmt.Println("[users]:", users)

	delete(users, 2) // removendo item do map
	fmt.Println("[users]:", users)

	usersInfo := map[int]map[string]string{
		1: {
			"name":     "Vinicius",
			"username": "Vilela",
		},
		2: {
			"name":     "John",
			"username": "Wick",
		},
	}
	fmt.Println("[usersInfo]:", usersInfo)

}
