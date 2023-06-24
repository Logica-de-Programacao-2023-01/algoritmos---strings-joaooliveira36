package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var palavra1, palavra2 string

	fmt.Print("Digite a primeira palavra: ")
	fmt.Scan(&palavra1)

	fmt.Print("Digite a segunda palavra: ")
	fmt.Scan(&palavra2)

	if Anagramas(palavra1, palavra2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func Anagramas(palavra1, palavra2 string) bool {
	if len(palavra1) != len(palavra2) {
		return false
	}

	letras_pal1 := strings.Split(palavra1, "")
	letras_pal2 := strings.Split(palavra2, "")

	sort.Strings(letras_pal1)
	sort.Strings(letras_pal2)

	return strings.Join(letras_pal1, "") == strings.Join(letras_pal2, "")
}
