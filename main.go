package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if !podeSacar {
		return "insuficient!", c.saldo
	}

	c.saldo -= valorDoSaque
	return "success!", c.saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if !podeDepositar {
		return "insuficient!", c.saldo
	}

	c.saldo += valorDoDeposito
	return "success!", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64, float64) {
	podeTransferir := c.saldo > 0 && valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo

	if !podeTransferir {
		return "insuficient!", c.saldo, contaDestino.saldo
	}

	c.saldo -= valorDaTransferencia
	contaDestino.saldo += valorDaTransferencia
	return "success!", c.saldo, contaDestino.saldo
}

func main() {
	contaDoMario := ContaCorrente{"Mario", 589, 123456, 125.50}
	contaDoFlavia := ContaCorrente{"Flavia", 589, 123456, 125.50}

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
