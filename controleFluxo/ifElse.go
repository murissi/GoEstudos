package main

import "fmt"

func main() {

	programaMenu()
	leComando()

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs..")
	// } else {
	// 	fmt.Println("Programa encerrado")
	// }

	comando := leComando()

	// Switch
	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo Logs...")

	case 0:
		fmt.Println("Saindo do programa")

	default:
		fmt.Println("Comando desconhecido")

	}

}

func programaMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

// funcao, retorna int
func leComando() int {
	var comandoLido int
	fmt.Print("Selecione uma opção: ")
	fmt.Scan(&comandoLido)
	return comandoLido
}
