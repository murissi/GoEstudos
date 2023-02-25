package main

import "fmt"

func main() {
	var nome string
	var ponteiro *string

	ponteiro = &nome
	*ponteiro = "Joao"

	fmt.Println(nome)
	fmt.Println(*ponteiro)
	fmt.Println()
	fmt.Println(&ponteiro)
	fmt.Println(&nome)

}
