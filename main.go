package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	// contaDoPedro := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Pedro", CPF: "124241414", Profissao: "Garoto de programa"}, NumeroAgencia: 1259, NumeroConta: 0001}
	// clientePedro := clientes.Titular{Nome: "Pedro", CPF: "124241414", Profissao: "Garoto de programa"}
	// contaDoPedro := contas.ContaCorrente{Titular: clientePedro, NumeroAgencia: 1259, NumeroConta: 0001, Saldo: 800}
	// contaDaMaria := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Maria", CPF: "12423333", Profissao: "Garota de programa"}, NumeroAgencia: 1258, NumeroConta: 0002}

	// status := contaDoPedro.Transferir(200, &contaDaMaria)
	// fmt.Println(status)
	// fmt.Println(contaDoPedro)
	// fmt.Println(contaDaMaria)
	// fmt.Println(contaDaMaria.ObterSaldo())

	contaDoPedro := contas.ContaPoupanca{}
	contaDoPedro.Depositar(100)
	PagarBoleto(&contaDoPedro, 60)
	fmt.Println(contaDoPedro.ObterSaldo())

	contaDaMaria := contas.ContaCorrente{}
	contaDaMaria.Depositar(500)
	PagarBoleto(&contaDaMaria, 400)
	fmt.Println(contaDaMaria.ObterSaldo())
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
