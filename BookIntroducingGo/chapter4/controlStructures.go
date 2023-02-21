package main

import "fmt"

func main() {

	// FOR
	for i := 1; i <= 10; i++ {

		for j := 10; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Print("Numero - 1 ao 3: ")
	var numeroSelecionado int
	fmt.Scan(&numeroSelecionado)

	switch numeroSelecionado {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Numero Invalido")
	}

	if numeroSelecionado%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}

	// Numero primo
	fmt.Print("Numero: ")
	var numeroPrimo int
	fmt.Scan(&numeroPrimo)

	var contador int
	for i := 1; i <= numeroPrimo; i++ {
		if numeroPrimo%i == 0 {
			contador++
		}
	}

	if contador == 2 {
		fmt.Printf("%v é numero primo\n", numeroPrimo)
	} else {
		fmt.Printf("%v não é um numero primo\n", numeroPrimo)
	}
}
