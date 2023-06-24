package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Print("Digite uma frase ou palavra: ")
	fmt.Scan(&palavra)
	palavra_maiuscula := strings.ToUpper(palavra)
	fmt.Print("A sua palavra em maiusculo é: ", palavra_maiuscula)
}
