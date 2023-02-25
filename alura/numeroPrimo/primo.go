package main

import "fmt"

func main() {
	fmt.Print("Digite numero inteiro: ")
	var numero int
	fmt.Scan(&numero)

	var contador int = 0
	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			contador++
		}
	}

	if contador == 2 {
		fmt.Printf("%v é Primo\n", numero)
	} else {
		fmt.Printf("%v não é Primo\n", numero)
	}
}
