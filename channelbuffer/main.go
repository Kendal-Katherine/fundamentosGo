package main

import "fmt"

//A diferença do channel normal para o buffer é que, o channel  com buffer tem um tamanho limitado de mensagens.
//E faz o envio e depois o recebimento, o channel normal, envia e recebe simutâneamente

func main() {
	//channel comum 	channel := make(chan int)
	//channel buffer 	channel := make(chan int, 100)
	channel := make(chan int)
	go setList(channel)

	//valor := <-channel //aqui lemos o que está no channel
	//mas podemos usar o for para ler  todos os valores do canal
	for v := range channel {
		fmt.Println("RECEBENDO: ", v)
	}
}

//a posição da seta <-  indica que estamos enviando um valor pro canal, ou seja, escrevendo em um canal
func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i //é dessa forma que escrevemos no channel
		fmt.Println("ENVIANDO: ", i)
	}
	close(channel) //fechamos o canal, assim sabendo que não ha mais nada a ser escrito nele
}
