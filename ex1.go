package main

import "fmt"

func main() {
	var x1, x2, x3 float64
	fmt.Print("Qual é o valor do primeiro número?")
	fmt.Scan(&x1)
	fmt.Print("Qual é valor do segundo número?")
	fmt.Scan(&x2)
	fmt.Print("Qual é valor do teceiro número?")
	fmt.Scan(&x3)
	soma := (x1 + x2 + x3)
	fmt.Println("A soma é", soma)

}
