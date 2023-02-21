package main

import "fmt"

func main() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scanf("%d", &comando) // Scanf recebe um valor
	/* Primeiro recebe o tipo da variavel, depois recebe & = endereco
	e logo em seguida a variavel */
	fmt.Println("Endereco da variavel comado", &comando)

	fmt.Println("Comando selecionado:", comando)

	fmt.Print("Idade: ")
	var idade int
	fmt.Scan(&idade) // Essa funcao é abreviacao da anterior
	// Porem sem a necessidade de indicar o tipo de variavel.
}
