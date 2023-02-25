package main

import "fmt"

func main() {
	// Built-in types: arrays, slices and maps

	// Arrays = sequencia de elementos do mesmo tipo.

	var y [5]int // 5 representa quantos elementos
	y[4] = 100
	fmt.Println(y)
	fmt.Println(y[4])

	// Arrays começam com o index 0.

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))

	// (_) = mostra pro go que nao precisamos disso.
	for _, value := range x {
		fmt.Printf("Valor: %v\n", value)
	}

	// Criando arrays de forma abreviada
	// lista := [5]float64{
	// 	98,
	// 	93,
	// 	77,
	// 	82,
	// 	83,
	// }

	// Slices
	// Igual a um array tem um tamanho mas é diferente por que é
	// permitido mudar o tamanho dele

	// x := make([]float64, 5, 10) // Um slice de tamanho 5 com capacidade para 10

	// Outra forma de criar slice
	// arr := [5]float64{1,2,3,4,5}
	// x := arr[0:5]

	// Append = adiciona um elemento no fim do slice
	slicez := []int64{1, 2, 3, 4, 5}
	slicez2 := append(slicez, 10, 20)
	fmt.Println(slicez, slicez2)

	// Copy
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) // criando slice com make([]int, 2) + tipo e tamanho.
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	/*
		Elementos do slice1 copiados para slice2, mas por slice2
		possuir apenas dois espaços. Não é possivel copiar tudo.
	*/

	// Slices são usadas para representar lista de itens.

	// MAP
	// Dicionarios, iguais do python

	dicionario := make(map[string]int)
	dicionario["key"] = 10
	fmt.Println(dicionario["key"])

	// Tipos de Erros
	// Compile-time errors = erros que acontecem na hora de compilar
	// runtime error = erros que acontecem na hora de rodar o programa

	p := make(map[int]int)
	x[1] = 10
	fmt.Println(p[1])

	// Deletando elementos
	delete(p, 1)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"]) // Retorna vazio " " que é valor padrao de String

	name, ok := elements["Un"]
	fmt.Println(name, ok) // 1. Retorna valor de procura
	// 2. Retorna se foi bem sucedido ou não.

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	var arrayNumeros [5]int
	arrayNumeros[4] = 800
	fmt.Println(arrayNumeros[4])

	pedaco := make([]int, 3, 9)
	fmt.Println(len(pedaco))

	arrayLetras := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(arrayLetras[2:5])

	sliceNumeros := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17, 2,
	}

	valorMenor := sliceNumeros[0]

	for i := 0; i < len(sliceNumeros); i++ {
		if sliceNumeros[i] < valorMenor {
			valorMenor = sliceNumeros[i]
		}
	}
	fmt.Println(valorMenor)

}
