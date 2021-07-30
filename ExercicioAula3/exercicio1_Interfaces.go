//Exercício #01
//Crie tipos para representar quadrados e círculos
//Crie uma interface que descreve o comportamento de calcular a área de uma forma geométrica com a seguinte assinatura: calculeArea() float64
//Implemente esse comportamento para os dois tipos criados
//Depois, crie uma função que tem como parâmetro a interface que você criou e que imprime o relatório do cálculo da área da forma geométrica
//Demonstre que seus tipos implementam a interface que você criou passando valores desses tipos como argumentos na chamada dessa função

package main

import (
	"fmt"
	"math"
)

type Forma interface {
	calculeArea() float64
}

type Quadrado struct {
	lado float64
}

func (q Quadrado) calculeArea() float64 {
	return math.Pow(q.lado, 2)
}

type Circulo struct {
	raio float64
}

func (c Circulo) calculeArea() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func reportarCalculo(forma Forma) {
	area := forma.calculeArea()
	fmt.Printf("A área desta forma geométrica é %.2f metros\n", area)
}

func main() {
	quadrado1 := Quadrado{lado: 5}
	circulo1 := Circulo{raio: 2.5}

	reportarCalculo(quadrado1)
	reportarCalculo(circulo1)
}
