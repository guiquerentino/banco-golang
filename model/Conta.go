package model

import (
	"fmt"
)

type Conta struct {
	NumeroConta string
	TipoConta   string
	SaldoConta  float64
}

func (p *Conta) TransfereSaldo(contaDestino Conta, valor float64) error {

	if p.SaldoConta < valor {
		fmt.Println("Saldo insuficiente!")
		return nil
	}

	p.SaldoConta -= valor
	contaDestino.SaldoConta += valor
	fmt.Println("Saldo transferido com sucesso!")
	return nil

}

func (p *Conta) AdicionaSaldo(contaDestino Conta, valor float64) {
	p.SaldoConta += valor
}

func (p *Conta) SacaSaldo(contaDestino Conta, valor float64) error {

	if p.SaldoConta < valor {
		fmt.Println("Saldo insuficiente!")
		return nil
	}
	p.SaldoConta -= valor
	return nil
}
