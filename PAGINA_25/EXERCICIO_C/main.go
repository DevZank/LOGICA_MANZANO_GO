/*
	Calcular e apresentar o valor do volume de uma lata de óleo, utilizando a fórmula:
	Volume * Raio * Altura
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var raioLata float64
	var alturaLata float64
	var volumeLata float64

	fmt.Print("Digite o Raio da lata: ")
	fmt.Scanln(&raioLata)
	fmt.Print("Digite a altura da lata: ")
	fmt.Scanln(&alturaLata)

	volumeLata = 3.14 * math.Pow(float64(raioLata), 2) * float64(alturaLata)

	fmt.Printf("O Volume da Lata é: %v", volumeLata)
}