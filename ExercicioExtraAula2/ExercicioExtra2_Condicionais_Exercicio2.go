//Exercício #02
//Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.

package main

import (
	"fmt"
)

func main() {
	A := 30

	if A == 0 {
		fmt.Println("A variável é o número zero.")
	} else if (A%2 == 0) && (A%3 == 0) {
		fmt.Printf("A variável %v é múltiplo de 2 e 3.", A)
	} else if A%2 == 0 {
		fmt.Printf("A variável %v é multiplo de 2.", A)
	} else if A%3 == 0 {
		fmt.Printf("A variável %v é multiplo de 3.", A)
	} else {
		fmt.Printf("A variável não é zero e nem múltiplo de 2 ou 3.")
	}
}
