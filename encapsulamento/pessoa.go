package encapsulamento

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereço
	DataDeNascimento time.Time
	Idade            int
}

//método - depois da palavra reservada func, abre e fecha (), com a variavel e a struct. Depois nome do método e o retorno
func (p *Pessoa) CalculaIdade() {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoDeNascimento
}

//função
/*func CalculaIdade(p Pessoa) int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
*/
