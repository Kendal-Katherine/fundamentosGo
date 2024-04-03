package main

import "fmt"

func main() {

	/* INICIANDO UM LAÇO
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	*/

	/* FOR COM LEN
	   texto := "palavra"
	   fmt.Println("Quantidade: ", len(texto) )
	   for i := 0; i < len(texto); i++ {
	   	fmt.Println(string(texto[i]))
	   	if string(texto[i]) == "r" {
	   		continue
	   	}
	   }
	*/

	/* FOR DENTRO DE OUTRO FOR*/
	for numBase := 1; numBase <= 10; numBase++ {
		for multiplicado := 1; multiplicado <= 10; multiplicado++ {
			fmt.Println(numBase, " x ", multiplicado, " = ", numBase*multiplicado)
		}
	}

}
