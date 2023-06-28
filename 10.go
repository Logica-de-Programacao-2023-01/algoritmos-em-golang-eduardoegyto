package main

import "fmt"

func main() {
	var pesoKilos float64
	fmt.Print("Digite o peso em quilos: ")
	fmt.Scanln(&pesoKilos)

	pesoLibras := pesoKilos * 2.20462
	fmt.Printf("O peso em libras é: %.2f\n", pesoLibras)
}
