package main

import "fmt"

func main (){

	lista := []int {2, 8, 3, 10, 5, 4, 7, 9, 1}
	
	numeroAte5 := 0
	numeroAte10 := 0

	for i := 0; i < len(lista); i++{
		if lista[i] <=5 {
			numeroAte5 += lista[i]
		} else{
			numeroAte10 += lista[i]
		}

	}

	fmt.Println(numeroAte5)
	fmt.Println(numeroAte10)
}