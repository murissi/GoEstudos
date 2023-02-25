package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	fmt.Println(leSitesDoArquivo())

}

func leSitesDoArquivo() []string {
	var sites []string

	arquivos, erros := os.Open("sites.txt")
	defer arquivos.Close()
	// nil == nullo(null).
	if erros != nil {
		fmt.Println("Ocorreu um erro: ", erros)
	}

	leitor := bufio.NewReader(arquivos)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			//fmt.Println(err)
			break
		}

	}

	return sites
}
