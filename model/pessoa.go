package model

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome       string
	Endereco   Endereco
	Nascimento time.Time
	Idade      int
}

func (p *Pessoa) CalculaIdade() {
	anoNascimento := p.Nascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoNascimento
	fmt.Println(p.Idade)
}
