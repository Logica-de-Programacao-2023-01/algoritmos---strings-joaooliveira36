package main

import "fmt"

func main() {
	var numero string
	ordem_crescente := true

	fmt.Print("Digite um núumero: ")
	fmt.Scan(&numero)

	for i := 1; i < len(numero); i++ {
		numero_atual := int(numero[i] - '0')
		numero_anterior := int(numero[i-1] - '0')
		if numero_anterior > numero_atual {
			ordem_crescente = false
			break
		}
	}
	if ordem_crescente {
		fmt.Print("Esta em ordem crescente")
	} else {
		fmt.Print("Não esta em ordem crescente")
	}
}
