package main

import "fmt"

//Função que recebe uma struct "Viagem" e retorna a Viagem mais cara utilizandos os valores da struct.

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	var viagemMaisCara Viagem

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		Viagem{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: "10/06/2023", Preco: 250.0},
		Viagem{Origem: "Rio de Janeiro", Destino: "Salvador", Data: "15/07/2023", Preco: 400.0},
		Viagem{Origem: "Salvador", Destino: "Fortaleza", Data: "20/08/2023", Preco: 300.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)

	fmt.Printf("Viagem mais cara:\nOrigem: %s\nDestino: %s\nData: %s\nPreço: R$ %.2f\n",
		viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Data, viagemMaisCara.Preco)
}
