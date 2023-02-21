package main

import "fmt"

func main() {

	fmt.Println("-----Verificar Sites-----")
	exibeMenu()
	respostaPrograma()

	fmt.Println("Digite sua idade: ")
	var idade int
	fmt.Scan(idade)

	fmt.Println(idade)

}

func exibeMenu() {
	fmt.Println("1 - Monitorar sites")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Fechar programa")
}

func respostaPrograma() int {
	fmt.Println("Seleciona uma opção valida: ")
	var resposta int
	fmt.Scan(resposta)
	return resposta
}
