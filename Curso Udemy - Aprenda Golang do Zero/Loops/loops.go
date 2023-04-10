package main

import (
	"fmt"
	"strings"
	"time"
)

func isPalidrome(str string) bool {
	var revesedString []string
	for _, char := range str {
		letter := string(char)
		revesedString = append([]string{letter}, revesedString...)
	}
	return str == strings.Join(revesedString, "")
}

func main() {

	fmt.Println("Analisador de palindromo")
	words := []string{"abba", "abc", "asda", "aabbaa", "ccaacc"}

	for i, word := range words {
		fmt.Printf("A palavra nº %d - \"%s\" é palindromo? %t \n", i, word, isPalidrome(word))
		time.Sleep(time.Second)
	}
}
