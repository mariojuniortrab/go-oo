package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoMario := ContaCorrente{"Mario", 589, 123456, 125.50}
	contaDaFlavia := ContaCorrente{"Flavia", 111, 113456, 125.0}

	var contaDaJoana *ContaCorrente
	contaDaJoana = new(ContaCorrente)
	contaDaJoana.titular = "Joana"

	fmt.Println(contaDoMario, contaDaFlavia, contaDaJoana)
}
