package main

import "fmt"

func main() {

	salario := 5001.00
	tipo := "PJ"

	if salario < 1.000 && tipo == "CLT" {
		salario -= (salario * 0.08)
	} else if salario <= 1000.00 && tipo == "PJ" {
		salario -= (salario * 0.5)
	} else if salario <= 1200.00 {
		salario -= (salario * 0.10)
	} else {
		salario -= (salario * 0.15)
	}

	fmt.Println("Salário: ", salario)
}

/*
	< - menor
	> - maior
	<= - menor igual
	>= - maior igual
*/
