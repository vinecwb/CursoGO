package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
