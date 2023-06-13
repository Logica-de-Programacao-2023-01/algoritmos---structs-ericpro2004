package main

import (
	"fmt"
)

//Função que recebe dois structs, de a partir dos dados de uma pessoa é possível descobrir onde ela mora.

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
	endereco := p.Endereco
	fmt.Printf("Endereço completo de %s:\n", p.Nome)
	fmt.Printf("Rua: %s, %d\n", endereco.Rua, endereco.Numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", endereco.Cidade, endereco.Estado)
}

func main() {
	endereco := Endereco{
		Rua:    "Rua Exemplo",
		Numero: 123,
		Cidade: "Cidade Exemplo",
		Estado: "Estado Exemplo",
	}
	pessoa := Pessoa{
		Nome:     "João",
		Idade:    30,
		Endereco: endereco,
	}
	imprimirEnderecoCompleto(pessoa)
}
