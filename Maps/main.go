package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["SP"] = 1000000
	m["CG"] = 9000000
	m["CG"] = 1000000
	m["RJ"] = 1000000
//DELETA UM ITEM DO MAP, MAS O VALOR DEVE SER IGUAL AO DO MAP, TENTEI "rj" E NÃO FUNCIONOU
	delete(m, "RJ")

//O LAÇO É DIFERENTE PARA O MAP
	for chave, valor := range m {
		fmt.Println("CIDADE: ", chave, "H: ", valor)
	}
/*

Estudar mais a questão do map
	valor, existe :=m["rj"]

	if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("chave não existe")
	}
*/
	fmt.Println(m)
}