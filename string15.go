package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	vogais := "aeiouAEIOU"
	nova_palavra := ""

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	for i := 0; i < len(palavra); i++ {
		caractere := string(palavra[i])
		if strings.Contains(vogais, caractere) {
			nova_palavra += "*"
		} else {
			nova_palavra += caractere
		}
	}
	fmt.Print(nova_palavra)
}
