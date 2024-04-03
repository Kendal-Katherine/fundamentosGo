package main

import "fmt"

func main()  {
	salario := 999.00
	var salarioMaisOBonus float64

	salarioMaisOBonus = salario

if salario <=  1000 {
	salarioMaisOBonus = (salarioMaisOBonus + 100) 
}
	fmt.Println("Salário: ", salarioMaisOBonus)
}
/*
	< - menor
	> - maior
	<= - menor igual
	>= - maior igual
*/