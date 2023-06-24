package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scan(&input)

	var uniqueLetters []string

	for _, char := range input {
		if strings.Count(input, string(char)) == 1 {
			uniqueLetters = append(uniqueLetters, string(char))
		}
	}

	uniqueString := strings.Join(uniqueLetters, "")
	fmt.Println("Letras únicas na string:", uniqueString)
}
