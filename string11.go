package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	tem_vogais := false
	vogais := "aeiouAEIOU"

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	for i := 0; i < len(palavra); i++ {
		if tem_vogais {
			break
		}
		if strings.Contains(vogais, string(palavra[i])) {
			fmt.Print("Sua string tem vogais")
			tem_vogais = true
		}
	}
}
