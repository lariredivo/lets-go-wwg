//Exercício Extra 2
//1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
//2) Como podemos pegar a informação do ano diretamente do console?

package main

import (
	"fmt"
	"time"
)

func main() {
	anoNascimento := 1996
	anoAtual := time.Now().Year()
	idade := anoAtual - anoNascimento

	fmt.Println(idade)

}
