package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra1, palavra2 string

	fmt.Print("Digite a primeira palavra: ")
	fmt.Scan(&palavra1)

	fmt.Print("Digite a segunda palavra: ")
	fmt.Scan(&palavra2)

	if strings.Contains(palavra1, palavra2) {
		fmt.Print("A segunda palavra é substring da primeira")
	} else {
		fmt.Print("A segunda palavra não é substring da primeira")
	}
}
