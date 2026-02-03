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
	fmt.Scan(&opecoes)
	fmt.Println("1 - soma")
	fmt.Println("2 - subtracao")
	fmt.Println("3 - multiplicacao")
	fmt.Println("4 - divisao")

	switch opcoes{
	case 1:
		fmt.Println(soma(x, y ))
	case 2:
		fmt.Println(subtracao(x y ))
	case 3:
		fmt.Println(multiplicacao(x y))
	case 4:
		fmt.Println(divisao(x, y))
	default:
		fmt.Println("Opçao inválida")	
	}

}
