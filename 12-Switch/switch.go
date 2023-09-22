package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Erro || O número deve ser de 1 à 7"
	}

}

// Outra forma de declarar
func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // Essa clausula joga para a próxima condição
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Erro || O número deve ser de 1 à 7"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(585)
	fmt.Println(dia)

	fmt.Println("-----------------------")
	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}
