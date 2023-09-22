package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2) //POR SER FORTEMENTE TIPADO NÃO ACEITA A SOMA DE TIPOS DIFERENTES
	fmt.Println(soma2)
	fmt.Println("--------------------------------------")

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("--------------------------------------")

	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println("--------------------------------------")

	//FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //AND
	fmt.Println(verdadeiro || falso) //OR
	fmt.Println(!verdadeiro)         // NEGAÇÃO
	fmt.Println(!falso)              //NEGAÇÃO
	fmt.Println("--------------------------------------")

	//FIM DOS RELACIONAIS

	// OPERADORES UNARIOS
	numero := 10
	numero++     // INCREMENTO
	numero += 15 // numero = numero + 15
	numero--     // DECREMENTO
	numero *= 3  // numero = numero * 3
	numero /= 10 // numero = numero / 10
	numero %= 3  // numero = numero % 3
	fmt.Println("--------------------------------------")

	fmt.Println(numero)
	// FIM DOS OPERADORES UNARIOS

	// FORMA DE USAR TERNÁRIO EM GO

	// texto := numero > 5 "Maior que 5" : "Menor que 5"  || ESPERADO EM OUTRAS LINGUAGENS
	// COMO FAZER EM GO \/

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
