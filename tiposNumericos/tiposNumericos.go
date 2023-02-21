// Tipos Numericos
package main

import (
	"fmt"
	"runtime"
)

type hotog int

func main() {
	// uint8  0 to 255
	// uint16 0 to 65.535
	// uint32 0 to 4.294.967.295
	// uint64 0 to 1.844.674.407.309.551.615

	// int8 -128 to 127
	// int16 -32768 to 32767
	// int32 -2.147.483.648 to 2.147.483.647
	// int64 -infinit to infinit

	// float32
	// float64

	// complex64
	//complex128

	//byte = uint8
	//rune = int32 (UTF8)

	// int e int32 não são a mesma coisa

	a := "e"
	b := "é"
	c := "#"

	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	x := 10
	y := 10.0
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	// Overflow
	/*
		Depois de chegar no limite e receber +1, o valor
		é recarregado e isso é um erro grave.
	*/
	var i uint16 = 65535
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)

	var v rune = 32
	fmt.Println(v)

}
