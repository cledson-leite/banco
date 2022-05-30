package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

// (conta *ContaCorrente) avisar que essa função usara qualquer variavel do tipo contacorrente como self
func (conta *ContaCorrente) Sacar(valor float64) string {
	if valor > 0 && valor <= conta.saldo {
		conta.saldo -= valor
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiente"
}
func (conta *ContaCorrente) Depositar(valor float64) string {
	if valor > 0 {
		conta.saldo += valor
		return "deposito realizado com sucesso"
	}
	return "Volres incorretos"
}

func main()  {
	contaDoCledson := ContaCorrente{titular: "Cledson", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	// contaDaRoberia := ContaCorrente{"Roberia", 222, 111222, 200}
	// o asterisco é um PONTEIRO para o local da memoria
	var contaDaCleide *ContaCorrente
	contaDaCleide = new(ContaCorrente)
	contaDaCleide.titular = "Cleide"
	contaDaCleide.saldo = 350.90

	fmt.Println(contaDoCledson.saldo)
	
	fmt.Println(contaDoCledson.Sacar(25.00))
	fmt.Println(contaDoCledson.saldo)

	fmt.Println(contaDoCledson.Sacar(250.00))
	fmt.Println(contaDoCledson.saldo)

	fmt.Println(contaDoCledson.Sacar(-25.00))
	fmt.Println(contaDoCledson.saldo)

	fmt.Println(contaDoCledson.Depositar(10))
	fmt.Println(contaDoCledson.saldo)
}