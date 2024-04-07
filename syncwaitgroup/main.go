package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go callDatabase(&wg)
	go callAPI(&wg)
	go processInternal(&wg)

	wg.Wait()

}

func callDatabase(wg *sync.WaitGroup) {
	//Simula  que estamos fazendo uma chamada a um  banco de dados
	time.Sleep(1 * time.Second) // mas o time sleep não funciona bem e precisamos usar o wg
	fmt.Println("Finalizado callDataBase")

	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	fmt.Println("Finalizado callAPI")
	time.Sleep(2 * time.Second)
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	fmt.Println("Finalizado processInternal")
	time.Sleep(1 * time.Second)
	wg.Done()
}
