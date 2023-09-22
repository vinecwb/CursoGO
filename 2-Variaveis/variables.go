package main

import "fmt"

// FORMAS PARA DECLARAÇÃO DE VARIAVEL NO GO
func main() {
	var variable1 string = "Variável 1" // DECLARAÇÃO DE TIPO DE VARIAVEL || FORMA EXPLICITA
	variable2 := "Variável 2"           // DECLARAÇÃO DE TIPO DE VARIAVEL || FORMA IMPLICITA
	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "lalala"
		variable4 string = "lalala"
	)

	fmt.Println(variable3, variable4)

	variable5, variable6 := "Variável 5", "Variável 6"
	fmt.Println(variable5, variable6)

	const constant1 string = "Constante 1"
	fmt.Println(constant1)

	variable5, variable6 = variable6, variable5 //INVERTENDO O VALOR DE DUAS VARIÁVEIS NO GO
	fmt.Println(variable5, variable6)

}
