package contas

import "banco/clientes"

type ContaCorrente struct {
	//quando os atributos do struct estao em minusculo ee fica privado, quanto a primeira maiuscula fica publico
	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	saldo float64
}

// (conta *ContaCorrente) avisar que essa função usara qualquer variavel do tipo contacorrente como self
func (conta *ContaCorrente) Saldo() float64 {
	return conta.saldo
}
func (conta *ContaCorrente) Sacar(valor float64) string {
	if valor > 0 && valor <= conta.saldo {
		conta.saldo -= valor
		return "Saque realizado com sucesso"
	}
	return "saldo insuficiente"
}

// (string, float64) range de retornos obrigatorios no caso de retornos multiplos
func (conta *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		conta.saldo += valor
		return "deposito realizado com sucesso", conta.saldo
	}
	return "Valores incorretos", conta.saldo
}
func (conta *ContaCorrente) Tranferencia(valor float64, destino *ContaCorrente) (string, float64, float64) {
	if valor > 0 && valor < conta.saldo {
		conta.Sacar(valor)
		destino.Depositar(valor)
		return "Transferência realizada com sucesso", conta.saldo, destino.saldo
	} else if valor < 0 {
		return "Valores incorretos", valor, conta.saldo
	}
	return "saldo insuficiente", valor, conta.saldo
}