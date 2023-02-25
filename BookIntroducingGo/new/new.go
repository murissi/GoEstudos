package main

import "fmt"

func sum(slice []int) int {
	var soma int
	for _, valores := range slice {
		soma += valores
	}
	return soma
}

func tipoNumero(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	numeros := []int{1, 2, 3, 4, 5}
	total := sum(numeros)
	fmt.Println(total)

	impar := tipoNumero(3)
	fmt.Println(impar)

}
