package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if !podeSacar {
		return "insuficient!", c.Saldo
	}

	c.Saldo -= valorDoSaque
	return "success!", c.Saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if !podeDepositar {
		return "insuficient!", c.Saldo
	}

	c.Saldo += valorDoDeposito
	return "success!", c.Saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64, float64) {
	podeTransferir := c.Saldo > 0 && valorDaTransferencia > 0 && valorDaTransferencia <= c.Saldo

	if !podeTransferir {
		return "insuficient!", c.Saldo, contaDestino.Saldo
	}

	c.Saldo -= valorDaTransferencia
	contaDestino.Saldo += valorDaTransferencia
	return "success!", c.Saldo, contaDestino.Saldo
}
