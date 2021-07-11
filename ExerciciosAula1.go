package main

import (
	"fmt"
)

func main() {
	// Exercício 1.
	//Utilizando a palavra reservada var declare uma variável numérica do tipo int.
	//Atribua um valor a essa variável.
	//Printe na tela o seu valor.
	//Repita para 3 variáveis diferentes.
	a := 27
	b := 401
	c := 10
	d := 15
	fmt.Println(a, b, c, d)

	// Exercicio 2.
	//Utilizando a marmota :=, declare duas variáveis: a e b.
	//Atribua os seguintes valores a elas, respectivamente: 230 e 27.
	//Crie variáveis para representar os resultados das operações matemáticas soma (a + b) e subtração (a - b).
	//Printe na tela os valores de todas as variáveis, um em cada linha.
	{
		a := 230
		b := 27
		sum := (a + b)
		sub := (a - b)
		fmt.Println(a, b, sum, sub)
	}

	//Exercicio 3
	//Calcule o valor total de uma compra que tem 3 itens representando o preço de todos eles em float64.
	// Todos os itens dessa compra precisams er comprados em mais de uma unidade.
	{
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

	//Declare duas variáveis do tipo string, uma vai receber o seu nome e outra a sua cor preferida.
	//Printe na tela uma frase utilizando os dois valores e o fmt.Printf() e o especificador de formato %s
	{
		nome := "Larissa"
		cor := "amarelo"

		fmt.Printf("Oi, meu nome é %s e minha cor favorita é %s", nome, cor)
	}

	{
		// Exercício 5 (Bool)
		//Declare 5 variáveis diferentes e atribua operações relacionais a elas.
		//Printe na tela uma linha para cada variável com o seu tipo e valor.
		a := 10 >= 1          //true
		b := 25 < 2           // false
		c := 10 == 18         // false
		d := 5 != 5           // false
		e := 89 > 0 && 50 > 0 //true

		fmt.Printf("o tipo de a é %T e seu valor é %v\n", a, a)
		fmt.Printf("o tipo de a é %T e seu valor é %v\n", b, b)
		fmt.Printf("o tipo de a é %T e seu valor é %v\n", c, c)
		fmt.Printf("o tipo de a é %T e seu valor é %v\n", d, d)
		fmt.Printf("o tipo de a é %T e seu valor é %v\n", e, e)
	}
	//Array - exemplo
	{
		var numeros = [4]int64{127, 290, 7, 83}
		fmt.Println(numeros[0]) //mostrar o primeiro número (posição 0)
	}

	// Escreva um programa que contenha um array de strings.
	// O valor de cada elemento deve ser o número do índice escrito por extenso.
	// Printe na tela o tipo do seu array e os valores de seus elementos.
	{
		var numeros [6]string
		numeros[0] = "zero"
		numeros[1] = "um"
		numeros[2] = "dois"
		numeros[3] = "três"
		numeros[4] = "quatro"
		numeros[5] = "cinco"

		fmt.Printf("o tipo é: %T\n", numeros)
		fmt.Println(numeros)
	}

	{
		// Escreva um programa que contenha um array de strings.
		// O valor de cada elemento deve ser o número do índice escrito por extenso.
		// Printe na tela o tipo do seu array e os valores de seus elementos.

		var numeros [6]string
		numeros[0] = "zero"
		numeros[1] = "um"
		numeros[2] = "dois"
		numeros[3] = "três"
		numeros[4] = "quatro"
		numeros[5] = "cinco"

		fmt.Printf("o tipo é: %T\n", numeros)
		fmt.Println(numeros)
	}

	{
		//Slice
		//Crie uma slice de string de tamanho 12 usando a função make() e atribua os signos do zodíaco a cada um dos elementos individualmente.
		//Printe na tela essa slice mostrando todos os seus elementos.
		//Printe também uma fatia dessa slice do índice 0 ao 2.

		signos := make([]string, 12)
		signos[0] = "aries"
		signos[1] = "touro"
		signos[2] = "gêmeos"
		signos[11] = "peixes"
		fmt.Println(signos)
		fmt.Println(signos[:2])
	}
	{
		//Maps = exemplo: dicionário
		//Crie um map chamado ano onde as chaves (key) sao os números dos meses do ano e o valor(value) é o nome que esse mês representa
		ano := make(map[int]string)
		ano[1] = "janeiro"
		ano[2] = "fevereiro"
		ano[3] = "março"
		ano[4] = "abril"
		ano[5] = "maio"
		ano[6] = "junho"
		ano[7] = "julho"
		ano[8] = "agosto"
		ano[9] = "setembro"
		ano[10] = "outubro"
		ano[11] = "novembro"
		ano[12] = "dezembro"

		fmt.Println(ano[1], ano[12])
		fmt.Println(len(ano))
	}
}
