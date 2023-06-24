package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	fmt.Print("Digite uma frase: ")
	fmt.Scan(frase)

	frase_atualizada := strings.ReplaceAll(frase, " ", "")
	fmt.Print(frase_atualizada)
}
