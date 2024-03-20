package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

// recebe valorDoSaque e retorna uma string
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	// valorDoSaque é maior que zero ou maior que o saldo
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) string {
	// podeDepositar := valorDeposito > 0 forma não refatorada de fazer
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Valor depositado na conta"
	} else {
		return "Valor inválido"
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
