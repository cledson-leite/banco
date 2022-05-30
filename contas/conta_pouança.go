package contas

import "banco/clientes"

type ContaPoupança struct {
	//quando os atributos do struct estao em minusculo ee fica privado, quanto a primeira maiuscula fica publico
	Titular clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo float64
}

// (conta *ContaPoupança) avisar que essa função usara qualquer variavel do tipo contaPoupança como self
func (conta *ContaPoupança) Saldo() float64 {
	return conta.saldo
}
func (conta *ContaPoupança) Sacar(valor float64) string {
	if valor > 0 && valor <= conta.saldo {
		conta.saldo -= valor
		return "Saque realizado com sucesso"
	}
	return "saldo insuficiente"
}

// (string, float64) range de retornos obrigatorios no caso de retornos multiplos
func (conta *ContaPoupança) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		conta.saldo += valor
		return "deposito realizado com sucesso", conta.saldo
	}
	return "Valores incorretos", conta.saldo
}
func (conta *ContaPoupança) Tranferencia(valor float64, destino *ContaPoupança) (string, float64, float64) {
	if valor > 0 && valor < conta.saldo {
		conta.Sacar(valor)
		destino.Depositar(valor)
		return "Transferência realizada com sucesso", conta.saldo, destino.saldo
		} else if valor < 0 {
			return "Valores incorretos", valor, conta.saldo
		}
		return "saldo insuficiente", valor, conta.saldo
	}

	func (conta *ContaPoupança) PagarBoleto(valor float64) (string, float64) {
		if valor > 0 {
			conta.Sacar(valor)
			return "deposito realizado com sucesso", conta.saldo
		}
		return "Valores incorretos", conta.saldo
	}