package main

import (
	"ex1/model"
	"fmt"
	"time"
)

func main() {
	var nomeItens []string
	nomeItens = append(nomeItens, "Arroz")
	nomeItens = append(nomeItens, "Feijao")
	nomeItens = append(nomeItens, "Macarrao")
	nomeItens = append(nomeItens, "Coca Cola")

	compra, err := model.NovaCompra("Mercado do João", time.Now(), nomeItens)

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(compra)
	}
}
