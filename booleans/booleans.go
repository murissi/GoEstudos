package main

import "fmt"

func main() {
	// Booleanos

	falso := false
	verdadeiro := true

	fmt.Printf("Falso: %v %T\nVerdadeiro: %v %T\n", falso, falso, verdadeiro, verdadeiro)
	z := (10 < 100)
	fmt.Println(z)

	i := 1 != 1
	fmt.Println(i)
}
