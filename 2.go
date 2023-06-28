package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	dobro := 2 * numero
	triplo := 3 * numero
	quadruplo := 4 * numero

	fmt.Println("O dobro é:", dobro)
	fmt.Println("O triplo é:", triplo)
	fmt.Println("O quadruplo é:", quadruplo)
}
