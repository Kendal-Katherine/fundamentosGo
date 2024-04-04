package main

import (
	"fmt"
	"fundamentosgo/Exer03/model"
	"time"
)

func main() {

	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "arroz")
	nomeDosItens = append(nomeDosItens, "feijão")
	nomeDosItens = append(nomeDosItens, "carne")
	nomeDosItens = append(nomeDosItens, "sabonete")

	compra, err := model.NewCompra("", time.Now(), nomeDosItens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}

}
