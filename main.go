package main

import (
	"fmt"
	"golangEstudos/model"
)

// func main() {
// 	fmt.Println("Hello, World!")

// 	endereco := model.Endereco{
// 		Rua:    "Rua 1",
// 		Numero: 123,
// 		Cidade: "Cidade 1",
// 	}

// 	pessoa := model.Pessoa{
// 		Nome:       "Dias",
// 		Endereco:   endereco,
// 		Nascimento: time.Date(2002, 10, 24, 0, 0, 0, 0, time.Local),
// 	}

// 	fmt.Println(pessoa)
// 	fmt.Println(endereco)

// 	pessoa.CalculaIdade()

// 	fmt.Println(pessoa.Idade)
// }

func main() {
	fmt.Println("Hello, World!")

	automovelMoto := model.Automovel{
		Ano:    2013,
		Placa:  "AXS-9F80",
		Modelo: "Ninja 300",
	}
	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 300,
	}

	fmt.Println(moto)
}
