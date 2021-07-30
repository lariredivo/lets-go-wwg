//Exercício #02
//Crie tipos que representam diferentes animais, com atributos que façam sentido para cada um deles
//Crie uma interface que descreve o comportamento de apresentar um animal com a seguinte assinatura: Apresenta()
//Cada animal saberá como se apresentar. Sendo assim, faça com que cada um dos tipos que você criou implemente o método Apresenta(), que deve printar uma frase apresentando o animal e seus atributos
//Demonstre que todos os tipos implementam a interface que você criou declarando uma slice de animais e percorrendo-a com um for range que, em todas as voltas, chama o método Apresenta().

package main

import (
	"fmt"
)

type apresentador interface {
	Apresenta()
}

type Cavalo struct {
	nome           string
	comidaFavorita string
}

func (c Cavalo) Apresenta() {
	fmt.Printf("Oi, meu nome é %s e minha comida favorita é %s\n", c.nome, c.comidaFavorita)
}

type Pombo struct {
	nome      string
	corDaPena string
}

func (p Pombo) Apresenta() {
	fmt.Printf("Oi, meu nome é %s e minhas penas são da cor %s\n", p.nome, p.corDaPena)
}

func main() {
	cavalo1 := Cavalo{
		nome:           "Alceu",
		comidaFavorita: "cenoura",
	}

	pombo1 := Pombo{
		nome:      "Chico",
		corDaPena: "cinza",
	}
	metazoa := []apresentador{cavalo1, pombo1}

	for _, v := range metazoa {
		v.Apresenta()
	}
}
