package main

import "fmt"

func main() {
	channel := make(chan int)
	go setList(channel)

	//valor := <-channel //aqui lemos o que está no channel
	//mas podemos usar o for para ler  todos os valores do canal
	for v := range channel {
		fmt.Println(v)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i //é dessa forma que escrevemos no channel
	}
	close(channel) //fechamos o canal, assim sabendo que não ha mais nada a ser escrito nele
}
