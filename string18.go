package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	numeros := "0123456789"

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	for i := 0; i < len(palavra); i++ {
		if strings.Contains(numeros, string(palavra[i])) == false {
			fmt.Print("Sua string nao tem somente numeros")
			return
		}
	}
	fmt.Print("Sua string so tem numeros")
}
