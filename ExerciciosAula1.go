package main

import (
	"fmt"
)

// Exercício 1.
//Utilizando a palavra reservada var declare uma variável numérica do tipo int.
//Atribua um valor a essa variável.
//Printe na tela o seu valor.
//Repita para 3 variáveis diferentes.
func main() {
	a := 27
	b := 401
	c := 10
	d := 15
	fmt.Println(a, b, c, d)
}

// Exercicio 2.
//Utilizando a marmota :=, declare duas variáveis: a e b.
//Atribua os seguintes valores a elas, respectivamente: 230 e 27.
//Crie variáveis para representar os resultados das operações matemáticas soma (a + b) e subtração (a - b).
//Printe na tela os valores de todas as variáveis, um em cada linha.
func main() {
	a := 230
	b := 27
	sum := (a + b)
	sub := (a - b)
	fmt.Println(a, b, sum, sub)
}

//Exercicio 3
//Calcule o valor total de uma compra que tem 3 itens representando o preço de todos eles em float64.
// Todos os itens dessa compra precisams er comprados em mais de uma unidade.
func main() {
	var pesoDoTomate, quantidadeDeGarrafasDeAzeite, unidadesDeSabonete float64
	pesoDoTomate = 2.600
	quantidadeDeGarrafasDeAzeite = 2
	unidadesDeSabonete = 7

	var precoDoTomate, precoDoAzeite, precoDoSabonete float64
	precoDoTomate = 10.3
	precoDoAzeite = 19
	precoDoSabonete = 5.80

	valorDaCompra := (pesoDoTomate * precoDoTomate) + (quantidadeDeGarrafasDeAzeite * precoDoAzeite) + (unidadesDeSabonete * precoDoSabonete)
	fmt.Println("O valor da compra deu %v, só uma garrafa de azeite já custou %v", valorDaCompra, precoDoAzeite)

}
