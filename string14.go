package main

import "fmt"

func main() {
	var numero string
	ordem_decrescente := true

	fmt.Print("Digite um núumero: ")
	fmt.Scan(&numero)

	for i := 1; i < len(numero); i++ {
		numero_atual := int(numero[i] - '0')
		numero_anterior := int(numero[i-1] - '0')
		if numero_anterior < numero_atual {
			ordem_decrescente = false
			break
		}
	}
	if ordem_decrescente {
		fmt.Print("Esta em ordem descrescente")
	} else {
		fmt.Print("Não esta em ordem descrescente")
	}
}
