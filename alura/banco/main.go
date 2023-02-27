package main

import "fmt"

// Classe
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// 0 value = valor inicial de uma variavel.

	// var contaDoGuilherme ContaCorrente = ContaCorrente{}
	contaDoGuilherme := ContaCorrente{
		titular: "Lucas",
		saldo:   125.5,
	}

	contaDaBruna := ContaCorrente{"Bruna", 222, 11122, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

}
