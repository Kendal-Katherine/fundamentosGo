package main

import "fmt"

func main() {
	var ehCarro = true
	var valorCarro = 1000.00
	
	if ehCarro {
		valorCarro += 55.50
	} else {
		valorCarro += 20.50
	}

	fmt.Println("Salário: ", valorCarro)
}

/*
	< - menor
	> - maior
	<= - menor igual
	>= - maior igual
*/
