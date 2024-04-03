package encapsulamento

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereço
	DataDeNascimento time.Time
}
//método - depois da palavra reservada func, abre e fecha (), com a variavel e a struct. Depois nome do método e o retorno
func (p Pessoa) IdadeAtual() int{
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}


//função
func CalculaIdade(p Pessoa) int{
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}