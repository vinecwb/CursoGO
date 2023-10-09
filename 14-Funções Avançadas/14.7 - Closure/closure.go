package main

import "fmt"

func closure() func() {
	texto := "Dentro da Função closure!"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da Função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
