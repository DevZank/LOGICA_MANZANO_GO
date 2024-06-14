/*
	Efetuar o cálculo e a apresentação do valor de uma prestação em atraso, utilizando a fórmula
	PRESTACAO = VALOR + (VALOR * TAXA/100) * TEMPO).
*/

package main

import (
	"fmt"
)

func main() {
	var valor float64
	var tempo int
	var taxa float64
	var total float64

	fmt.Print("Qual o valor das prestações? ")
	fmt.Scanln(&valor)
	fmt.Print("Qual a quantidade de dias atrasados? ")
	fmt.Scanln(&tempo)
	fmt.Print("Qual a taxa de Juros? ")
	fmt.Scanln(&taxa)

	total = valor + valor * taxa / 100 * float64(tempo)

	fmt.Println("")
	fmt.Printf("O valor final da prestação é R$%v", total)
}