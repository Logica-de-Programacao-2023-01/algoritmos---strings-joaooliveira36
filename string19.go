package main

import (
	"fmt"
)

func main() {
	var palavra string
	var palavra_ao_contrario string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	for i := len(palavra) - 1; i >= 0; i-- {
		palavra_ao_contrario += string(palavra[i])
	}
	fmt.Print(palavra_ao_contrario)
}
