package main

import (
	"fmt"
	"os"
)

func ReadFile() {

	//tem a finalidade de continuar a execução do código, mesmo com o panic
	defer func() {//função anônima
		if r := recover(); r != nil{
			fmt.Println("Recuperado")
		}
	}()


	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("FileNotExist")//Com o panic o código é interrompido aqui
	}
}

func main() {
	
	ReadFile()

	fmt.Println("Fim")
}
