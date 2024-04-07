package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	for i := 0; i < 1000; i++ {
		//o go realiza uma função assincrona
		go showMessage(strconv.Itoa(i))

	}
	//espera a função executar
	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))

}

func showMessage(message string) {
	fmt.Println("This is message: ", message)
}
