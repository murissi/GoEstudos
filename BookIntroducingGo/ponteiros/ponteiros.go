package main

import "fmt"

type Tarefa struct {
	titulo  string
	proxima *Tarefa
}

func main() {
	t1 := Tarefa{titulo: "Comprar frutas"}
	t2 := Tarefa{titulo: "comprar amaciante"}
	t3 := Tarefa{titulo: "gravar video "}

	t1.proxima = &t2
	t2.proxima = &t3
	// lista := t1

	idade := 31
	var p *int

	p = &idade

	fmt.Println(idade) // valor idade
	fmt.Println()
	fmt.Println(p)      // variavel p aponta para idade
	fmt.Println(&idade) // endereco de idade
	fmt.Println(&p)     // endereco de p
	fmt.Println(*p)     // mostra o valor que esta guardado na variavel que p aponta

	*p = 59
	idade = 59

	fmt.Printf("%v \n%v\n", *p, idade)
}
