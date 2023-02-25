package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World")
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {

	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0

	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))

	// average(xs)
	average(xs)
	helloWorld()

	// Variadic Functions
	// Varios parametros numa funcao

	fmt.Println(add(1, 2, 3, 4, 5))

	slicee := []int{1, 2, 3, 4} // array
	// sliceee := [4]int{4, 3, 2, 1} // array
	fmt.Println(slicee)

	slicee = append(slicee, 5, 6)
	fmt.Println(slicee)

	// Closure
	// Posivel criar funcoes dentro de funcoes
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	// Quando criamos uma funcao local, ela pode ser acessada por
	// outras variaveis locais.
	o := 0
	increment := func() int {
		o++
		return o
	}

	increment()
	increment()
	fmt.Println(o)

	// Recursion = Chama ela mesmo

}
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))

}

// Funcoes nome + parametro + tipo de entrada + retorno.
// func average(xs []float64) float64 {
// 	panic("Not implemented")
// }
