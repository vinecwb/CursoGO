package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // o & é necessário para verificar o Endereço de memória

	fmt.Println(variavel3, ponteiro)  //Aqui irá apresentar o endereço de memória
	fmt.Println(variavel3, *ponteiro) // Aqui nós desreferenciamos o endereço e solicitamos a exibição do valor

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) //Valor do ponteiro vai mudar
	fmt.Println(variavel3, ponteiro)  // Porém o endereço continuará o mesmo
}
