package contas

type ContaCorrente struct {
	//quando os atributos do struct estao em minusculo ee fica privado, quanto a primeira maiuscula fica publico
	Titular string
	NumeroAgencia int
	NumeroConta int
	Saldo float64
}

// (conta *ContaCorrente) avisar que essa função usara qualquer variavel do tipo contacorrente como self
func (conta *ContaCorrente) Sacar(valor float64) string {
	if valor > 0 && valor <= conta.Saldo {
		conta.Saldo -= valor
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiente"
}

// (string, float64) range de retornos obrigatorios no caso de retornos multiplos
func (conta *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		conta.Saldo += valor
		return "deposito realizado com sucesso", conta.Saldo
	}
	return "Valores incorretos", conta.Saldo
}
func (conta *ContaCorrente) Tranferencia(valor float64, destino *ContaCorrente) (string, float64, float64) {
	if valor > 0 && valor < conta.Saldo {
		conta.Sacar(valor)
		destino.Depositar(valor)
		return "Transferência realizada com sucesso", conta.Saldo, destino.Saldo
	} else if valor < 0 {
		return "Valores incorretos", valor, conta.Saldo
	}
	return "Saldo insuficiente", valor, conta.Saldo
}