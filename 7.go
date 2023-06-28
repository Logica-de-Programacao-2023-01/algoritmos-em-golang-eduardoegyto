package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	sucessor := numero + 1
	antecessor := numero - 1

	fmt.Println("O sucessor é:", sucessor)
	fmt.Println("O antecessor é:", antecessor)
}
