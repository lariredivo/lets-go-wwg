//Exercício #02
//Considere um cenário em que você precise regularmente trabalhar com slices de inteiros e extrair a soma e média dos números dessa lista. Usando o que você aprendeu sobre métodos, qual seria sua solução para esse problema em Go?

package main

import (
	"fmt"
)

type conjunto []int

func (c conjunto) Some() int {
	soma := 0
	for _, operando := range c {
		soma += operando
	}
	return soma
}

func (c conjunto) CalculeMedia() float64 {
	soma := float64(c.Some())
	quantidade := float64(len(c))
	return soma / quantidade

}

func main() {
	conjunto1 := conjunto{3, 7, 27, 40}

	soma1 := conjunto1.Some()
	media1 := conjunto1.CalculeMedia()
	fmt.Printf("Nosso conjunto tem os elementos: %v\n\tsoma: %d\n\tmédia: %f", conjunto1, soma1, media1)

}
