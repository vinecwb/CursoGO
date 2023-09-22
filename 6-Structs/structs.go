package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Vinicius"
	u.idade = 29
	fmt.Println(u)

	// SEGUNDA FORMA DE DECLARAR \/
	usuario2 := usuario{"Davi", 21}
	fmt.Println(usuario2)

	// TERCEIRA FORMA DE DECLARAR \/ CASO ONDE VOCÊ NÃO TENHA ALGUMA INFO DO STRUCT
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
