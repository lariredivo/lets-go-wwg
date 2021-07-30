//Exercício #02
//Construa uma função que receba uma palavra ou frase e uma letra, e retorne o número de ocorrências da letra informada.
package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Minha ideologia é: floresta em pé, água limpa, ar puro e comida sem veneno"
	letra := "a"

	var final = strings.Count(texto, letra)
	fmt.Printf("A letra %s foi encontrada %d vezes na frase '%s'", letra, final, texto)
}
