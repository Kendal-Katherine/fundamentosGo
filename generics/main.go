package main

import "fmt"

func main() {

	slice1 := []int{0, 1, 2, 3}
	slice2 := []string{"a", "b", "c", "d"}
	newInts := reverse(slice1)
	newString := reverse(slice2)
	fmt.Println(newInts)
	fmt.Println(newString)

}

//Um único método genérico que resolve o problema de inverter um slice qualquer tipo
func reverse[T int | string](slice []T) []T {

	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}

//Outra forma de usar métodos genéricos  é criar uma interface com os tipos necessários 
//e usar ela como  parâmetro do método genérico.
type constraintCustom interface{
	int | string
}

func reverse2[T constraintCustom](slice []T) []T {

	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}
