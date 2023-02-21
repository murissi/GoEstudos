package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	// x := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem."
	// y := `oi bom dia\ncomo vai?\tespero \"que\" tudo bem.` // nenhum codigo é interpretado usando ``

	// fmt.Println(x, y)

	// FMT
	// x := "oi"
	// y := "bom dia"

	// z := fmt.Sprint(x, " ", y) // concatena e retorna string
	// fmt.Println(z)

	x := 10
	x = int(b) // Cast
	fmt.Printf("%v\n", x)

	fmt.Printf("%v\n", b)

	a := 10.10
	b := int(a)

	fmt.Println(b)

}
