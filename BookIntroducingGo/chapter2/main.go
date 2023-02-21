package main

import "fmt"

func main() {
	fmt.Println("50 é multiplo de 5: ", 50%5)

	// Strings
	a := len("Hello, World") // Mostra tamanho de texto
	fmt.Println(a)

	b := "Cachorro"
	fmt.Println(b[0]) // Mostra a letra de acordo com a tabela utf-8

	// Concatenacao
	x := "Hello"
	fmt.Println(x + "Mundo")

	fmt.Println(len("Hello"))
	fmt.Println("Hello, World"[1]) // Pegando um byte, que é um inteiro
	fmt.Println("Hello, " + "World")

	// Booleans
	/*
		&& and
		|| or
		! not
	*/

	fmt.Println(true && true)                                            // t
	fmt.Println(true && false)                                           // f
	fmt.Println(true || true)                                            // t
	fmt.Println(true || false)                                           // t
	fmt.Println(!true)                                                   // f
	fmt.Println((true && false) || (false && true) || !(false && false)) // t

}

// Go is statically typed programming language. This means
// that variables always have a specific type and that type cannot change
