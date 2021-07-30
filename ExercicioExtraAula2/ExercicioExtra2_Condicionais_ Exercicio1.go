//Exercício #01 (https://womenwhogocwb.gitbook.io/letsgo/condicionais/exercicios-extras)
//Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor

package main

import (
	"fmt"
)

func main() {
	A := 10
	B := 48
	C := 345

	if (A > B) && (A > C) {
		fmt.Printf("A variável com o maior valor é: A com o valor %v", A)
	} else if (B > A) && (B > C) {
		fmt.Printf("A variável com o maior valor é: B com o valor %v", B)
	} else if (C > A) && (C > B) {
		fmt.Printf("A variável com o maior valor é: C com o valor %v", C)
	}
}
