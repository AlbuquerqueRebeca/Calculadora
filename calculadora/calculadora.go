package main

import "fmt"

func main() {

	var x, y int
	var opcoes int

	fmt.Println("Digite o número: ")
	fmt.Scan(&x)
	fmt.Println("Digite o número:")
	fmt.Scan(&y)

	fmt.Println("Escolha a operação: ")
	fmt.Println("1 - soma")
	fmt.Println("2 - subtracao")
	fmt.Println("3 - multiplicacao")
	fmt.Println("4 - divisao")

}
