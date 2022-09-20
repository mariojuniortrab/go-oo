package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteMario := clientes.Titular{Nome: "Mario", CPF: "123.123.123.11", Profissao: "Programador"}
	contaDoMario := contas.ContaCorrente{Titular: clienteMario, NumeroAgencia: 589, NumeroConta: 123456}
	contaDoMario.Depositar(100)

	fmt.Println(contaDoMario.ObterSaldo())

}
