package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if !podeSacar {
		return "insuficient!", c.saldo
	}

	c.saldo -= valorDoSaque
	return "success!", c.saldo
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if !podeDepositar {
		return "insuficient!", c.saldo
	}

	c.saldo += valorDoDeposito
	return "success!", c.saldo
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) (string, float64, float64) {
	podeTransferir := c.saldo > 0 && valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo

	if !podeTransferir {
		return "insuficient!", c.saldo, contaDestino.saldo
	}

	c.saldo -= valorDaTransferencia
	contaDestino.saldo += valorDaTransferencia
	return "success!", c.saldo, contaDestino.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
