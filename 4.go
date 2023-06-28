package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	mediaPonderada := (2*num1 + 3*num2 + 5*num3) / 10
	fmt.Printf("A média ponderada é: %.2f\n", mediaPonderada)
}
