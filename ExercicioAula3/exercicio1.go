//Exercício #01
//Construa dois métodos: um que retorna a área e outro que retorna o perímetro de uma estrutura que representa um círculo. Printe a área e o perímetro de um círculo.

package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func (c Circulo) CalculeArea() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c Circulo) CalculePerimetro() float64 {
	return 2 * math.Pi * c.raio
}

func main() {
	circulo1 := Circulo{
		raio: 5,
	}

	areaDoCirculo := circulo1.CalculeArea()
	perimetroDoCirculo := circulo1.CalculePerimetro()
	fmt.Printf("círculo:\n\traio: %.4f\n\tárea: %.4f\n\tperímetro: %.4f", circulo1.raio, areaDoCirculo, perimetroDoCirculo)
}
