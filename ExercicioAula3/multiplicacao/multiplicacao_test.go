package multiplicacao

import "testing"

func TestMultiplique(t *testing.T) {
		testes := []struct{
			nome string

			x int

			y int

			esperado int
		}{
			nome: "7 multiplicado por 5 é igual a 35",
			x: 7,
			y: 5,
			esperado: 35,
		},
		{
			nome: "3 multiplicado por 20 é igual a 60",
			x: 3,
			y: 20,
			esperado: 60,
		},
		{
			nome: "13 multiplicado por 1000 é 13000"
			x: 13,
			y: 1000,
			esperado: 13000,
		},
		{
			nome: "7 multiplicado por 3 é 21"
			x: 3,
			y:	21,
			esperado: 14,
		},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T){
			obtive := Multiplique(teste.x, teste.y)
			
			if obtive != teste.esperado {
				t.Errorf(form:"espero '%d'; mas obtive '%d'", teste.esperado, obtive)
			}
		})
	}

	t.Run(name:"7 multiplicado por 5 é igual a 35", func(t *testing.T) {
	//chamar a função que multiplica
	obtive := Multiplique(x:7, y:5)
	// dizer o que esperamos que aconteça
	espero := 35
	// comparar o que recebemos com o que esperávamos receber
	if obtive != espero {
		//sinalize que o teste falhou
		t.Errof(format:"espero '%d', mas obtive '%d'", espero, obtive)
	}
})
	t.Run(name:"3 multiplicado por 20 é igual a 60", func(t *testing.T){
	obtive := Multiplique(x: 3, y:20)
	espero := 60

	if obtive != espero {
		t.Errorf(format:"espero '%d', mas obtive '%d'", espero, obtive)
	}
})
