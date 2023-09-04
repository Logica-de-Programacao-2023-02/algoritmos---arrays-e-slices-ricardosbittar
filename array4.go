package main

import "fmt"

func main() {
	var slice []int
	var x, y, z int
	var soma int
	for {
		fmt.Println("Digite o primeiro número: ")
		fmt.Scanln(&x)
		fmt.Println("Digite o segundo número: ")
		fmt.Scanln(&y)
		fmt.Println("Digite o terceiro número: ")
		fmt.Scanln(&z)

		if x < 0 || y < 0 || z < 0 {
			fmt.Println("nenhum numero pode ser negativo neste caso")
		} else {
			break
		}
	}

	slice = append(slice, x, y, z)

	for i := 0; i < len(slice); i++ {

		soma += slice[i]
	}

	fmt.Println("A soma dos valores é", soma)
}
