package main

import "fmt"

// var x int = 10 // Variavel de nivel package scope
var a int
var b float64
var c string
var d bool

func main() {
	// x := 10.0 // operador de declaração
	// y := "Olá, bom dia"

	// fmt.Printf("x: %v, %T\n", x, x)
	// fmt.Printf("y: %v, %T\n", y, y)

	// x = 20.0 // operador de atribuição
	// fmt.Printf("x: %v, %T\n", x, x)

	// var (
	// 	a, b, c string = "Olá", "Bem", "Vindo"
	// 	d, e, f int    = 1, 2, 3
	// )
	// fmt.Println(a, b, c, d, e, f)

	// qualquerCoisa(4)

	//x := true // declaraco
	//x = false // atribuicao

	// Valor padrao
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}

func qualquerCoisa(x int) {
	fmt.Println(x)
}
