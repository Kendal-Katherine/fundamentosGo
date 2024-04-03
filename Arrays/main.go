package main

import "fmt"

func main() {
	/*ARRAY NÃO PODE SER MUDADO*/
	var listaToda = [10]int{2, 10, 9, 4, 5, 8, 1, 3}
	//listaToda = append(listaToda, 1)

	fmt.Printf("%T", listaToda)
	
}
