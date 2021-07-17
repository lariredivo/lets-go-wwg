package main

import (
	"fmt"
)

//String - Exercício 7 (https://womenwhogocwb.gitbook.io/letsgo/tipos-compostos/exercicios)
//Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
//Crie três variáveis do tipo pessoa.
//Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.
type Pessoa struct {
	nome         string
	idade        int
	corPreferida string
}

func main() {
	pessoa1 := Pessoa{
		nome:         "Larissa",
		idade:        24,
		corPreferida: "amarela",
	}

	pessoa2 := Pessoa{
		nome:         "Lisa",
		idade:        47,
		corPreferida: "azul",
	}

	pessoa3 := Pessoa{
		nome:         "Marcela",
		idade:        19,
		corPreferida: "roxo",
	}

	fmt.Println(pessoa1, pessoa2, pessoa3)
	fmt.Printf("%s: tem %d anos e sua cor preferida é %s\n", pessoa1.nome, pessoa1.idade, pessoa1.corPreferida)
	fmt.Printf("%s: tem %d anos e sua cor preferida é %s\n", pessoa2.nome, pessoa2.idade, pessoa2.corPreferida)
	fmt.Printf("%s: tem %d anos e sua cor preferida é %s\n", pessoa3.nome, pessoa3.idade, pessoa3.corPreferida)
}

//Condicionais - If - Exercício 1 (https://womenwhogocwb.gitbook.io/letsgo/condicionais/exercicios)
//Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
//Declare uma variável e atribua a ela o ano atual.
//Escreva um programa que verifique se no ano atual essa pessoa já está apta a votar e que printe na tela uma mensagem informando caso positivo e caso negativo.

func main() {
	anoDeNascimento := 1973
	anoAtual := 2021

	idade := anoAtual - anoDeNascimento

	if idade >= 16 {
		fmt.Printf("Você tem %d anos e já pode votar!", idade)
	} else {
		fmt.Printf("Você tem %d anos e ainda não pode votar!", idade)
	}
}

// Exercício 2
// Declara uma variável;
//Atribui o valor de um número a ela;
//Informa se o número é positivo ou negativo.

func main() {
	numero := -24

	if numero < 0 {
		fmt.Println("Negativo")
	} else numero > 0 {
		fmt.Println("Positivo")
	} else {
		fmt.Println("O número é zero!")
	}
}

//Exercicio 3
//Faça um programa que, dada a idade de uma pessoa, verifique se ela é menor de idade, maior de idade ou idosa.
func main() {
	idade := 24

	if idade < 18 {
		fmt.Println("Menor de idade")
	} else if idade <= 60 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Idoso")
	}
}

//Switch -> similar ao if mas com melhor visibilidade
func main() {
	idade := 40
	
	switch {
	case idade >=60:
		fmt.Println("Idoso")
	case idade >=18 && idade < 60:
		fmt.Println("Maior de idade")
	case idade <0:
		fmt.Println ("O valor informado não é uma idade valida")
	default:
		fmt.Println("Menor de idade")
	}
}

//Declarar uma variável chamada hora e atribuir um valor int a ela.
//Usando switch elencar cases e printar na tela se a hora corresponde ao período da manhã, tarde, noite ou madrugada

func main() {
	hora := 12

	switch {
	case hora >= 0 && hora < 6:
		fmt.Println("Madrugada")
	case hora > 6 && hora <= 12:
		fmt.Println("Manhã")
	case hora >=12 && hora < 18:
		fmt.Println("Tarde")
	case hora >= 18 && hora < 24:
		fmt.Println("Noite")
	default:
		fmt.Println("Valor informado não é válido")
	}
	//Laços de repetição (https://womenwhogocwb.gitbook.io/letsgo/lacos-de-repeticao/conteudo)
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

//Exercicio 1 - Utilizando a estrutura for, printe na tela os números inteiros de 13 a 27 (inclusive).(https://womenwhogocwb.gitbook.io/letsgo/lacos-de-repeticao/exercicios)
func main() {
	for i := 13; i <= 27; i++ {
		fmt.Println(i)
	}
}

//outra forma do exemplo de cima (similar ao while)
func main() {
	i := 13
	for i <= 27 {
		fmt.Println(i)
		i++
	}	
}

//Exercicio 2 - Utilizando uma sintaxe da estrutura for diferente da usada no exercício acima, printe na tela todas as horas do dia (usando o formato de 24 horas).
func main() {
	hora := 0
	for hora <= 24 {
		fmt.Printf("%.2d:00\n", hora)
		hora++
	}	
}

//Exercicio 3 - Utilizando a resolução anterior, demonstre também todos os minutos de um dia (formato 24 horas).
func main() {
	hora := 0
	for hora < 24 {
		for minuto := 0; minuto < 60; minuto++ {
			fmt.Printf("%.2d:%.2d\n", hora, minuto)
		}
		hora++
	}
}

//Usando a função "for", printe na tela tofos os segundos de 00:00:00 até 02:59:59
func main() {
	hora := 0
	for hora < 3 {
		for minuto := 0; minuto < 60; minuto++ {
			for segundo := 0; segundo < 60; segundo++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundo)
				}
		}
		hora++
	}
}

//for range
func main() {
	var fatia = []string{"zero", "um", "dois", "três", "quatro"}
	for indice, value := range fatia {
		fmt.Println(indice, value)
	}
}

//Dada uma slice de string que representa uma lista de mercado, use o for range para percorrê-la e printe na tela cada um dos itens de cmpra ao lado da sua ordem de inserção na lista
func main() {
	var listaDeCompras = []string{"abacate", "sabonete", "azeite", "tomate", "banana"}
	for indice, value := range listaDeCompras { 
		fmt.Printf("%d) %s\n", indice+1, value)
	}
}
