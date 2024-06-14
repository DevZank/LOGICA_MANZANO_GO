/*
	Ler uma temperatura em graus Celsius e apresentá-la convertida em graus Fahrenheit. A fórmula de
	conversão é F = (9 * C + 160) / 5, sendo F a temperatura em Fahrenheit e C a temperatura em Celsius
*/

package main

import (
	"fmt"
)

func main() {
	var celsius float32

	fmt.Print("Escreva uma Temperatura em ºC: ")
	fmt.Scanln(&celsius)

	var fahrenheit float32 = (9 * celsius + 160) / 5

	fmt.Println("=================================")
	fmt.Printf("%vºC convertido para Fahrenheit é: %vºF\n", celsius, fahrenheit)
	fmt.Println("=================================")
}