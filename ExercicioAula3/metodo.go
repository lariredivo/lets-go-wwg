// https://womenwhogocwb.gitbook.io/letsgo/metodos/conteudo

package main

import (
	"fmt"
)

type Gato struct {
	nome                  string
	pessoaResponsavel     string
	saborDeSachePreferido string
}

func (g Gato) Apresenta() {
	fmt.Printf("Miau! Meu nome é %s, a pessoa responsável por mim é %s e meu sabor favorito de sache é %s\n", g.nome, g.pessoaResponsavel, g.saborDeSachePreferido)

}

type Cachorro struct {
	nome               string
	pessoaResponsavel  string
	brinquedoPreferido string
}

func (c Cachorro) Apresenta() {
	fmt.Printf("AuAu! Meu nome é %s, a pessoa responsável por mim é %s e meu brinquedo favorito é %s\n", c.nome, c.pessoaResponsavel, c.brinquedoPreferido)
}

func main() {
	gato1 := Gato{
		nome:                  "Mingau",
		pessoaResponsavel:     "Shakira",
		saborDeSachePreferido: "frango",
	}

	gato2 := Gato{
		nome:                  "Lua",
		pessoaResponsavel:     "Anitta",
		saborDeSachePreferido: "carne",
	}

	cachorro1 := Cachorro{
		nome:               "Belinha",
		pessoaResponsavel:  "Thalia",
		brinquedoPreferido: "bolinha",
	}

	gato1.Apresenta()
	gato2.Apresenta()
	cachorro1.Apresenta()
}
