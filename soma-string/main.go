package main

import (
	"fmt"
	"reflect"
)

func main(){
	nome := "Kendal"
	b := " Grazi"
	result := nome + b

	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

}