package main

import "fmt"

func main() {
	var listaToda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	segundaLista := listaToda[:3]
	terceiraLista := listaToda[4:]
	ultimoItem := listaToda[len(listaToda)-1:]

	fmt.Println("Segunda lista", segundaLista)
	fmt.Println("Terceira lista", terceiraLista)
	fmt.Println("Último item da lista", ultimoItem)

}
