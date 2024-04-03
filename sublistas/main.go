package main

import "fmt"

func main() {
	var listaToda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	var listaMenorQueCinco = make([]int, 0)
	fmt.Println(listaMenorQueCinco)

	for i := 0; i < len(listaToda); i++ {
		if listaToda[i] < 9 {
			listaMenorQueCinco = append(listaMenorQueCinco, listaToda[i])
			fmt.Println("Posição [", i , "]", listaMenorQueCinco)
		}
	}

	fmt.Println(listaMenorQueCinco)

}
