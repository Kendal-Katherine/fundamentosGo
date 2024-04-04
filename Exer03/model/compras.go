package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado    string
	DataCompra time.Time
	itens      []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NewCompra(mercado string, date time.Time, nomeDosItens []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("Mercado é obrigatório")
	}

	var itens []ItemDaCompra

	for _, nome := range nomeDosItens {

		itens = append(itens, ItemDaCompra{Nome: nome})
	}

	return &Compra{
		Mercado:    mercado,
		DataCompra: time.Now(),
		itens:      itens,
	}, nil

}
