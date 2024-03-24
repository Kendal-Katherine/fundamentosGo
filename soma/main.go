package main

import (
	"fmt"
	"reflect"
)

func main(){
	a := 10.0
	b := 20.0
	result := a + b

	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

}