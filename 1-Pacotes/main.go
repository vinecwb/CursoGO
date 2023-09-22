package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	assistant.Escrever()

	//Pacote externo que valida se um email é real
	erro := checkmail.ValidateFormat("teste_vsd_tes@gmail.com")
	fmt.Println(erro)
}
