package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	clienteMario := clientes.Titular{Nome: "Mario", CPF: "123.123.123.11", Profissao: "Programador"}
	contaDoMario := contas.ContaCorrente{Titular: clienteMario, NumeroAgencia: 589, NumeroConta: 123456}
	contaDoMario.Depositar(100)
	contaDoMarioPoupanca := contas.ContaPoupanca{Titular: clienteMario, NumeroAgencia: 589, NumeroConta: 123456}
	contaDoMarioPoupanca.Depositar(100)

	PagarBoleto(&contaDoMario, 20)
	PagarBoleto(&contaDoMarioPoupanca, 30)

	fmt.Println(contaDoMario.ObterSaldo())
	fmt.Println(contaDoMarioPoupanca.ObterSaldo())

}
