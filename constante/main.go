package main

import ("fmt"
"reflect"
)

func main() {
	const num = 10 
	// num = 5 não pode ser subscrevido

	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))
}