/*
	Ler uma temperatura em graus Fahrenheit e apresentá-la convertida em graus Celsius. A fórmula de
	conversão é C = (F - 32) * (5/9) , sendo F a temperatura em Fahrenheit e C a temperatura em Celsius.
*/

package main

import (
	"fmt"
)

func main() {
	var fahrenheit float32

	fmt.Print("Escreva uma temperatura em ºF: ")
	fmt.Scanln(&fahrenheit)

	var celsius float32 = (fahrenheit - 32) / 1.8

	fmt.Println("==================================")
	fmt.Printf("%vºF convertido para Celsius é: %vºC\n", fahrenheit, celsius)
	fmt.Println("==================================")
}