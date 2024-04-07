package main

import "os"

func main() {

	_, err := os.Open("f:/settings.txt")

	if err != nil{
		panic(err)//panic dá um shutdown na aplicação e não é usado para tudot
	}

}