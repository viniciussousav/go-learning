package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	var email string = "vinicius@gmail.com"

	err := checkmail.ValidateFormat(email)

	if err != nil {
		fmt.Printf("Email %s is invalid", email)
		return
	}

	fmt.Printf("Email \"%s\" is valid", email)

}
