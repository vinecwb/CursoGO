package main

import (
	"errors"
	"fmt"
)

func main() {
	// TIPOS DE INTEIROS NO GO: int (Vai usar a arquitetura como base),int8, int16, int32, int64

	var numero int16 = 10000
	fmt.Println(numero)

	numero1 := 12345041
	fmt.Println(numero1)

	//UINT // UNSYGNED INT - Int sem sinal
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = INT 8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12321321.45
	fmt.Println(numeroReal2)

	//PEGOU O 64 DA ARQUITETURA DO COMPUTADOR
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NUMEROS REAIS

	var str string = "Qualquer Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B' //ENTRE ASPAS SIMPLES ELE VAI RETORNAR O NÚMERO DO CARACTER NA TABELA ASCII
	fmt.Println(char)

	// FIM STRINGS

	//VALOR ZERO || QUANDO SE INICIA UMA VARIAVEL SEM INSERIR UM VALOR A ELA
	var texto int
	fmt.Println(texto)

	var boolean bool = true
	fmt.Println(boolean)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
