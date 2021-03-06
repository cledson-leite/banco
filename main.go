package main

import (
	"fmt"
	//caminho para importar packege personalisados
	"banco/clientes"
	"banco/contas"
)

func PagarBoleto(conta ContaInterface, valor float64) {
	conta.Sacar(valor)
}

type ContaInterface interface {
	Sacar(valor float64) string
}

func main()  {
	cledson := clientes.Titular{Nome: "Cledson", CPF: "1234568901", Profissao: "dev"}
	roberia := clientes.Titular{Nome: "Roberia", CPF: "1098654321", Profissao: "dev"}
	contaDoCledson  := contas.ContaCorrente{Titular: cledson , NumeroAgencia: 589, NumeroConta: 123456}
	contaDoCledson.Depositar(1000)
	contaDaRoberia := contas.ContaCorrente{Titular: roberia, NumeroAgencia: 222, NumeroConta: 111222}
	contaDaRoberia.Depositar(1000)

	// o asterisco é um PONTEIRO para o local da memoria
	var contaDaCleide *contas.ContaPoupança
	contaDaCleide = new(contas.ContaPoupança)
	contaDaCleide.Titular.Nome = "Cleide"
	contaDaCleide.Depositar(350.90)

	fmt.Println(contaDoCledson.Saldo())
	
	fmt.Println(contaDoCledson.Sacar(25.00))
	fmt.Println(contaDoCledson.Saldo())

	fmt.Println(contaDoCledson.Sacar(250.00))
	fmt.Println(contaDoCledson.Saldo())

	fmt.Println(contaDoCledson.Sacar(-25.00))
	fmt.Println(contaDoCledson.Saldo())

	status, valor := contaDoCledson.Depositar(10)

	fmt.Println(status,valor)
	
	fmt.Println(contaDoCledson.Saldo(), contaDaRoberia.Saldo())
	fmt.Println(contaDoCledson.Tranferencia(50, &contaDaRoberia))
	
	fmt.Println(contaDoCledson.Saldo(), contaDaRoberia.Saldo())
	fmt.Println(contaDoCledson.Tranferencia(500, &contaDaRoberia))
	
	fmt.Println(contaDoCledson.Saldo(), contaDaRoberia.Saldo())
	fmt.Println(contaDoCledson.Tranferencia(-50, &contaDaRoberia))
	
	fmt.Println(contaDoCledson.Saldo())
	PagarBoleto(&contaDoCledson, 100)
	fmt.Println(contaDoCledson.Saldo())
	
	fmt.Println(contaDaCleide.Saldo())
	PagarBoleto(contaDaCleide, 100)
	fmt.Println(contaDaCleide.Saldo())

}