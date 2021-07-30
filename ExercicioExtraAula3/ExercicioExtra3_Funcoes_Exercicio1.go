//Exercício #01
//Construa uma função que receba uma lista de números inteiros, modifique essa lista dobrando os números ímpares e divida por dois os pares, e retorne a lista modificada e a soma de todos os elementos da lista.

package main

import (
	"fmt"
)

func main() {
	var numeros = []int{14, 28, 45, 3, 9, 13, 68, 5}
	var total int = 0

	for i, n := range numeros {
		switch {
		case (n % 2) != 0:
			numeros[i] *= 2

		case (n % 2) == 0:
			numeros[i] /= 2
		}
		total += numeros[i]
	}
	fmt.Printf("%v\n%d", numeros, total)
}
