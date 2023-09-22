package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // Dessa forma o escopo é local do IF
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que zero")
	} else {
		fmt.Println("Número é menor que menos dez")
	}

}
