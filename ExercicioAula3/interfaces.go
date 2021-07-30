// https://womenwhogocwb.gitbook.io/letsgo/interfaces/conteudo exemplo do conteúdo Interfaces
package main

import (
	"fmt"
)

type Apresentador interface {
	Apresenta()
}

type Humano struct {
	nome         string
	corPreferida string
}

func (h Humano) Apresenta() {
	fmt.Printf("Olá! Meu nome é %s e minha cor preferida é %s\n", h.nome, h.corPreferida)
}

type Cachorro struct {
	nome               string
	brinquedoPreferido string
}

func (c Cachorro) Apresenta() {
	fmt.Printf("Auau! Meu nome é %s e meu brinquedo prefeirdo é %s\n", c.nome, c.brinquedoPreferido)
}

type Gato struct {
	nome                  string
	saborPreferidoDoSache string
}

func (g Gato) Apresenta() {
	fmt.Printf("Miau! Meu nome é %s e meu sabor de sache preferido é %s\n", g.nome, g.saborPreferidoDoSache)
}

type Calopsita struct {
	nome            string
	tamanhoDoTopete string
}

func (s Calopsita) Apresenta() {
	fmt.Printf("Croc! Meu nome é %s e meu topete tem o tamanho de %s", s.nome, s.tamanhoDoTopete)
}

func main() {
	humano1 := Humano{
		nome:         "Shakira",
		corPreferida: "laranja",
	}
	cachorro1 := Cachorro{
		nome:               "Champion",
		brinquedoPreferido: "bolinha",
	}
	gato1 := Gato{
		nome:                  "Fifi",
		saborPreferidoDoSache: "cordeiro",
	}

	gato2 := Gato{
		nome:                  "Mingau",
		saborPreferidoDoSache: "frango",
	}

	calopsita1 := Calopsita{
		nome:            "Calo",
		tamanhoDoTopete: "médio",
	}

	apresentadores := []Apresentador{humano1, cachorro1, gato1, gato2, calopsita1}

	for _, apresentador := range apresentadores {
		apresentador.Apresenta()
	}
}
