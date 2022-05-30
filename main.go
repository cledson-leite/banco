package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func main()  {
	contaDoCledson := ContaCorrente{titular: "Cledson", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	contaDaRoberia := ContaCorrente{"Roberia", 222, 111222, 200}
	// o asterisco é umPONTEIRO para o local da memoria
	var contaDaCleide *ContaCorrente
	contaDaCleide = new(ContaCorrente)
	contaDaCleide.titular = "Cleide"
	contaDaCleide.saldo = 350.90

	fmt.Println(contaDoCledson, contaDaRoberia)
	fmt.Println(contaDaCleide, *contaDaCleide)
}