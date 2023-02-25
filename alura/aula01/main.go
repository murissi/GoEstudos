package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func exibeIntroducao() {
	nome := "Lucas"
	versao := 1.1

	fmt.Printf("Olá, sr(a). %v\n", nome)
	fmt.Printf("Esse programa está na versão %v\n", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

}
func leComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Printf("O valor da variável comando é %v\n", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando....")
	site := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
		"https://www.youtube.com.br",
	}

	for _, valores := range site {
		respota, _ := http.Get(valores)
		if respota.StatusCode == 200 {
			fmt.Println("Site:", valores, "Foi carregado com sucesso!\n")
			fmt.Println()
			time.Sleep(5 * time.Second)

		} else {
			fmt.Println("Site", valores, " está com problemas. Error:", respota.StatusCode, ".\n")
			time.Sleep(5 * time.Second)

		}

	}
}

//	func devolveNomeIdade() (string, int) {
//		nome := "lucas"
//		idade := 22
//		return nome, idade
//	}
func main() {

	// For sem nada == while
	exibeIntroducao()
	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

		// nome, _ := devolveNomeIdade()
		// fmt.Println(nome)
	}
}
