package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero string

	fmt.Print("Digite um número: ")
	fmt.Scan(numero)

	teste_num, err := strconv.ParseFloat(numero, 64)

	if err != nil {
		fmt.Printf("A entrada %s não e um float", numero)
		return
	} else {
		fmt.Printf("A entrada %s é um float em %s", numero, teste_num)
	}
}
