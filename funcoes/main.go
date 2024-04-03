package main

import "fmt"

//Função Init inicia sempre em primeiro lugar não precisa ser chamada no método main
func init() {
	fmt.Println("Função init")
}

func main() {
	//Executando Função
	ImprimirMensagem()
	//Executando Função com Parâmetros
	ImprimirMensagemCP("Mensagem X")
	ImprimirMensagemCP("Mensagem Y")
	//Executando Função com Parâmetros e lógica
	ImprimirMensagemCPL("Kendal")
	ImprimirMensagemCPL("Katherine")
	//Executando Função com retorno
	resultado := Soma(10, 15)
	fmt.Println(resultado)
	resultado2 := Soma(15, 20)
	fmt.Println(resultado2)
	//Função com mais de um retorno
	soma, divisao, subtracao, multiplicacao := Operacoes(5, 5)
	fmt.Println(soma, divisao, subtracao, multiplicacao)
	//Função com retornos nomeados
	soma2, divisao2, subtracao2, multiplicacao2 := OperacoesNomeadas(3, 3)
	fmt.Println(soma2, divisao2, subtracao2, multiplicacao2)

}

/*------------------------------------------------------------------------------------------------------------*/

//Função sem parâmetros
func ImprimirMensagem() {
	fmt.Println("Essa é uma função")
}

//Função com parâmetros
func ImprimirMensagemCP(mensagem string) {
	fmt.Println(mensagem)
}

//Função com parâmetros e lógica
func ImprimirMensagemCPL(mensagem string) {
	mensagem += ", bom dia!"
	fmt.Println(mensagem)
}

//Função com retorno
func Soma(numero1 int, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}

//Função com retorno de mais de um valor
func Operacoes(numero1 int, numero2 int) (int, int, int, int) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	divisao := numero1 / numero2
	multiplicacao := numero1 * numero2
	return soma, subtracao, divisao, multiplicacao
}

//Função com retorno nomeado
func OperacoesNomeadas(numero1 int, numero2 int) (soma int, subtracao int, divisao int, multiplicacao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	divisao = numero1 / numero2
	multiplicacao = numero1 * numero2
	return
}
