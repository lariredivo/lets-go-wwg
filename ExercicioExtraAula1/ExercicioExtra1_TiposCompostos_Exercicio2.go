//Exercício #02
//Considerando os times do item anterior, crie uma slice para representar cada um.
//Adicione o jogador Luis Inácio no time vermelho usando a função append()
//Printe na tela os nomes dos jogadores do time vermelho

package main

import (
	"fmt"
)

func main() {

	timeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}
	timeVermelho = append(timeVermelho, "Luis Inácio")

	fmt.Println(timeVermelho)
}
