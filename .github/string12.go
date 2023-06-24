package main

import (
	"fmt"
)

func main() {
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	if Palindromo_check(palavra) {
		fmt.Println("A palavra é um palíndromo.")
	} else {
		fmt.Println("A palavra não é um palíndromo.")
	}
}

func Palindromo_check(palavra string) bool {

	for i := 0; i < len(palavra)/2; i++ {
		if palavra[i] != palavra[len(palavra)-i-1] {
			return false
		}
	}

	return true
}
