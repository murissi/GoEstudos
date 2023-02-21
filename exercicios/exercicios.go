// %v = mostra o valor %T = mostra tipo

package main

import "fmt"

type texto string

func main() {

	count := 4

	if count == 4 {
		fmt.Printf("O valor informado foi: %v\n", count)
	}

	switch count {
	case 4:
		fmt.Println("Valor informado foi 1")
	case 2:
		fmt.Println("Valor informado foi 2")
	default:
		fmt.Println("Valor informado foi 3")

	}
}
