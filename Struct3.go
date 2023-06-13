package main

import "fmt"

//Função que recebe um struct de um triângulo e retorna a área do mesmo.

type Triangulo struct {
	Base   float64
	Altura float64
}

func CA(t Triangulo) float64 {
	return t.Base * t.Altura / 2
}

func main() {
	triangulo := Triangulo{Base: 5.0, Altura: 3.0}
	area := CA(triangulo)
	fmt.Printf("Área do triângulo: %.2f\n", area)
}
