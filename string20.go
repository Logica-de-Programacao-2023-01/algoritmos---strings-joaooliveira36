package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scan(&input)

	firstRune := []rune(input)[0]
	isCamelCase := unicode.IsUpper(firstRune)

	wordCount := 1
	for _, char := range input {
		if unicode.IsUpper(char) {
			wordCount++
		}
	}

	if isCamelCase {
		fmt.Println("A string está em camelCase.")
	} else {
		fmt.Println("A string não está em camelCase.")
	}
	fmt.Println("Quantidade de palavras:", wordCount)
}
