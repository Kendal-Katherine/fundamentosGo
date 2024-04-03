package main

import "fmt"

func main() {
	x := 5
	y := &x //aqui não estamos passando o valor e sim a referência de memória de X AULA 39
	*y = 10 //Toda vez que for usar uma varíavel de ponteiro tem que colocar o *, x e y passam a ter o mesmo valor, pq y era a memória de x, e acabamos de substituir o valor da memória
	fmt.Println("Main----------------")
	fmt.Println(x, *y)
	fmt.Println(&x, y)
	ImprimirValores(&x, y)
}

func ImprimirValores(x *int, y *int) {
	fmt.Println("Imprimir Valores----------------")
	fmt.Println(x, y)
	
}
