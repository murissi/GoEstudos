package main

import (
	"fmt"
)

func main() {
	// Strings
	/*
		%v %T
		Raw string literals
		Conversão para slice of bytes: []byte(x)
		%#U, %#x

	*/

	// Um texto é um slice de bytes.

	s := "Hello, playground"
	// mostrar(s)

	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v\n%T\n - %#U - %#x\n", v, v, v, v)

	}

}

func mostrar(texto string) {

	fmt.Println(texto)

}
