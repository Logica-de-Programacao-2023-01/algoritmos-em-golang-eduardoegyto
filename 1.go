package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&x)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&y)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&z)

	soma := x + y + z
	fmt.Println("A soma dos números é:", soma)
}
