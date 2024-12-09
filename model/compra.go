package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemCompra
}

type ItemCompra struct {
	Nome string
}

func NovaCompra(mercado string, data time.Time, nomeItens []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("Mercado não pode ser vazio")
	}

	if len(nomeItens) == 0 {
		return nil, errors.New("Itens não podem ser vazios")
	}

	var itens []ItemCompra

	for _, nome := range nomeItens {
		itens = append(itens, ItemCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    time.Now(),
		Itens:   itens,
	}, nil
}
