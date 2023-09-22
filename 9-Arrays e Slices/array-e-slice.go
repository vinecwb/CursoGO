package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)
	// Segunda opção para popular

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	fmt.Println(slice)

	slice = append(slice, 22) // vai pegar todas as informações do slice anterior inserindo o valor a mais
	fmt.Println(slice)
	fmt.Println("-------------------------------")
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)
	fmt.Println("-----------------------")
	//Arrays Internos
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacidade

}
