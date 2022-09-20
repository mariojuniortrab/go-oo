package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
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

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
