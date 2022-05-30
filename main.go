package main

import (
	"fmt"

	//caminho para importar packege personalisados
	"banco/contas"
)



func main()  {
	contaDoCledson := contas.ContaCorrente{Titular: "Cledson", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.50}
	contaDaRoberia := contas.ContaCorrente{"Roberia", 222, 111222, 200}
	// o asterisco é um PONTEIRO para o local da memoria
	var contaDaCleide *contas.ContaCorrente
	contaDaCleide = new(contas.ContaCorrente)
	contaDaCleide.Titular = "Cleide"
	contaDaCleide.Saldo = 350.90

	fmt.Println(contaDoCledson.Saldo)
	
	fmt.Println(contaDoCledson.Sacar(25.00))
	fmt.Println(contaDoCledson.Saldo)

	fmt.Println(contaDoCledson.Sacar(250.00))
	fmt.Println(contaDoCledson.Saldo)

	fmt.Println(contaDoCledson.Sacar(-25.00))
	fmt.Println(contaDoCledson.Saldo)

	status, valor := contaDoCledson.Depositar(10)

	fmt.Println(status,valor)
	
	fmt.Println(contaDoCledson.Saldo, contaDaRoberia.Saldo)
	fmt.Println(contaDoCledson.Tranferencia(50, &contaDaRoberia))
	
	fmt.Println(contaDoCledson.Saldo, contaDaRoberia.Saldo)
	fmt.Println(contaDoCledson.Tranferencia(500, &contaDaRoberia))
	
	fmt.Println(contaDoCledson.Saldo, contaDaRoberia.Saldo)
	fmt.Println(contaDoCledson.Tranferencia(-50, &contaDaRoberia))

}