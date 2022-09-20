package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoMario := contas.ContaCorrente{Titular: "Mario", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.50}
	contaDoFlavia := contas.ContaCorrente{Titular: "Flavia", NumeroAgencia: 589, NumeroConta: 123457, Saldo: 225.50}

	status, valor := contaDoMario.Sacar(30)
	fmt.Println(status, valor)
	status, valor = contaDoMario.Sacar(300)
	fmt.Println(status, valor)
	status, valor = contaDoMario.Depositar(500)
	fmt.Println(status, valor)

	status, valor1, valor2 := contaDoMario.Transferir(100, &contaDoFlavia)
	fmt.Println(status, valor1, valor2)

	status, valor1, valor2 = contaDoFlavia.Transferir(125, &contaDoMario)
	fmt.Println(status, valor1, valor2)

}
