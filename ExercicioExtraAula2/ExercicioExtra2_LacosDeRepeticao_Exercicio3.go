//Exercício #03
//Faça um programa que, dado um texto inserido pelo usuário, itere nesse texto e conte o número de ocorrências de cada letra. Em seguida imprima em ordem alfabética a letra e o número de ocorrências dela no texto informado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Minha ideologia é: floresta em pé, água limpa, ar puro e comida sem veneno"
	letra := "a"

	var final = strings.Count(texto, letra)
	fmt.Printf("A letra %s foi encontrada %d vezes na frase %s", letra, final, texto)
}
