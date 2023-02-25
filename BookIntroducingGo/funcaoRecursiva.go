package main

import "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fatorial(n-1)
	}
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {

	fmt.Println(fatorial(5))

	// Defer, panic and recover
	// defer programa para uma funcao ser lida depois de outras funcoes
	defer second()
	first()

	// defer é bom para usar em funcoes que esquecemos tipo sc.close() no Java.

	// Panic and Recover
	// Panic cria um runtime error

	// panic("PANIC")
	// str := recover()
	// fmt.Println(str)

	// Pointers
	/*
		Localizacao na memoria onde
		 um valor é armazenado do que ele mesmo
	*/

	x := 5
	zero(&x)
	fmt.Println(x)
}
