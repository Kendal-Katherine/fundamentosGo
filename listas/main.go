package main

import "fmt"

func main() {
/*
	lista := []int{4, 9, 6, 7}
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista posição 1: ", lista[0])
	fmt.Println("Lista posição 2: ", lista[1])
	fmt.Println("Lista posição 3: ", lista[2])
	fmt.Println("Lista tamanho: ", len(lista))
	fmt.Println()
	fmt.Println("Depois do append")
	lista = append(lista, 8)
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista tamanho: ", len(lista))
*/

/*USANDO A FUNÇÃO MAKE*/

lista := make([]int, 1)
lista[0] = 5
lista = append(lista, 2)
lista = append(lista, 3)

fmt.Printf("%T", lista)

somaTotal := 0

for i := 0; i < len(lista); i++{
	somaTotal +=lista[i]
	fmt.Println(lista[i])
}

fmt.Println("Média: ", somaTotal/len(lista))

}