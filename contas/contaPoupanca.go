package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                                 clientes.Titular
	NomeumeroAgencia, NumeroConta, Operacao int
	saldo                                   float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	// valorDoSaque é maior que zero ou maior que o saldo
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) string {
	// podeDepositar := valorDeposito > 0 forma não refatorada de fazer
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Valor depositado na conta"
	} else {
		return "Valor inválido"
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
