package main

import "fmt"

func main() {

	count := 0
	soma := 0
	for {

		fmt.Print("Digite um numero: ")
		var numero int
		fmt.Scan(&numero)

		count++
		soma += numero

		fmt.Print("Continuar | S | N | ")
		var decisao string
		fmt.Scan(&decisao)

		if decisao == "N" {
			break
		} else {
			continue
		}

	}

	fmt.Println("Resultado: ", soma/count)

}
