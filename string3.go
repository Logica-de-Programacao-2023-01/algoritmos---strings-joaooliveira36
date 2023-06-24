package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	var letra_antiga string
	var letra_nova string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(palavra)

	fmt.Print("Qual letra voce quer substituir: ")
	fmt.Scan(letra_antiga)

	fmt.Print("Qual letra voce quer colocar no lugar da(s) antiga(s): ")
	fmt.Scan(letra_nova)

	palavra_atualizada := strings.ReplaceAll(palavra, letra_antiga, letra_nova)
	fmt.Print(palavra_atualizada)
}
