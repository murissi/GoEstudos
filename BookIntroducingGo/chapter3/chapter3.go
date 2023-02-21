package main

import "fmt"

func main() {

	x := "First"
	fmt.Println(x)

	x += " Second"
	fmt.Println(x)

	a := "Hello"
	b := "World"
	fmt.Println(a == b)

	b = "Hello"
	fmt.Println(a == b)

	// Escopo = Corpo de uma funcao {} ou ()
	/*

		Variaveis criadas a nivel package
		podem ser acessadas por todas as
		 funcoes.

	*/

	// Constantes
	const constante string = "Hello, World"
	fmt.Println(constante)

	const constSemTipo = 1.333
	fmt.Println(constSemTipo)

	// Definindo multiplas variaveis = pode usar var() ou const()
	// var (
	// 	c = 44
	// 	d = 45
	// 	f = 46
	// )

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input) // Lê saida de teclado.

	output := input * 2

	fmt.Println(output)

	// Maneiras de criar variaveis
	// var novaVariavel int == usada no package scope
	// variavelNova :=  == mais usada
	// var NovaNova =

	// Escopo é o espaço criado em uma funcao. { } ou ()

	// Variavel pode ser alterada, const nao

	fmt.Print("Digite valor em Fahreiheit: ")
	var valorFah float64
	fmt.Scanf("%f", &valorFah)

	calculoCelsius := (valorFah - 32) * 5 / 9.0
	fmt.Println(calculoCelsius)

	fmt.Print("Digite metros: ")
	var peses float64
	fmt.Scanf("%f", &peses)

	pesesCalculados := peses * 0.3048

	fmt.Println(pesesCalculados)

}
