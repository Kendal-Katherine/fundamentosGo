package main

import (
	"fmt"
	"fundamentosgo/encapsulamento"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := encapsulamento.Endereço{
		Rua:    "Rua X",
		Numero: 15,
		Cidade: "São Paulo",
	}
	pessoa := encapsulamento.Pessoa{
		Nome:             "Kendal",
		Endereco:         endereco,
		DataDeNascimento: time.Date(1997, 8, 30, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(pessoa)
	fmt.Println(endereco)
	//idade := encapsulamento.CalculaIdade(pessoa)
	pessoa.CalculaIdade() // o método faz parte da struct de pessoa agora e é mais elegante
	fmt.Println(pessoa.Idade)

	endereco.Numero = 18
	fmt.Println(endereco.Numero)

	/******************************************************************************/
	automovelMoto := encapsulamento.Automovel{
		Ano: 2022,
		Placa: "XPTO",
		Modelo: "CG",
	}

	moto := encapsulamento.Moto{
		Automovel: automovelMoto,
		Cilindradas: 125,
	}
	fmt.Println(moto)
	fmt.Println(moto.Placa)

	
}
