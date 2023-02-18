package main

import (
	"fmt"
	"reflect" // biblioteca que descobre o tipo da variavel
)

//

func main() {
	// ; no final não é indicado.
	fmt.Println("Hello World!!!!")

	// Declarando variavel
	var nome string = "Lucas"
	var versao float32 = 1.1
	var idade int = 22
	// variaveis sem nada, sempre iniciam ou com 0 ou 0.0 ou vazio.
	// Go da erro se a variavel foi declarada mas nao foi usada.

	var idadeSemTipo = 24
	// Go consegue descobrir que tipo é variavel

	fmt.Println("Olá, sr.", nome, "Sua idade é: ", idade)
	fmt.Println("Este program está na versão.", versao)
	fmt.Println(idadeSemTipo)
	fmt.Println("O tipo da variavel nome é: ", reflect.TypeOf(nome))
	fmt.Println("Tipo varivel idade: ", reflect.TypeOf(idade))

	// Operador de variaveis curto
	// Não precisa usar var e nem declarar o tipo
	cpf := "222.222.222-22"
	fmt.Print(cpf)

}
