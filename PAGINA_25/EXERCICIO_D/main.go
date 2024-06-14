/*
	Efetuar o cálculo da quantidade de litros de combustível gasta em uma viagem, utilizando um
	automóvel que faz 12 Km por litro. Para obter o cálculo, o usuário deve fornecer o tempo gasto
	(TEMPO) e a velocidade média (VELOCIDADE) durante a viagem. Desta forma, será possível obter a
	distância percorrida com a fórmula DISTANCIA = TEMPO * VELOCIDADE. Possuindo o valor da
	distância, basta calcular a quantidade de litros de combustível utilizada na viagem com a fórmula
	LITROS_USADOS = DISTANCIA / 12. Ao final, o programa deve apresentar os valores da velocidade
	média (VELOCIDADE), tempo gasto na viagem (TEMPO), a distancia percorrida (DISTANCIA) e a
	quantidade de litros (LITROS_USADOS) utilizada na viagem.
*/

package main

import (
	"fmt"
)

func main() {
	var tempo int
	var velocidade float64
	var distancia float64
	var litros_usados float64

	fmt.Print("Qual o tempo gasto da Viagem? ")
	fmt.Scanln(&tempo)
	fmt.Print("Qual a velocidade media percorrida na Viagem? ")
	fmt.Scanln(&velocidade)

	distancia = float64(tempo) * velocidade
	litros_usados = distancia / 12

	fmt.Println("------------------------------------------")
	fmt.Printf("Nesta Viagem, o tempo gasto foi de: %vh\n", tempo)
	fmt.Printf("A velocidade média foi de: %vkm/h\n", velocidade)
	fmt.Printf("A distância foi de: %vkm\n", distancia)
	fmt.Printf("A quantidade de litros gastos foi de: %.2fL\n", litros_usados)
	fmt.Println("------------------------------------------")
}