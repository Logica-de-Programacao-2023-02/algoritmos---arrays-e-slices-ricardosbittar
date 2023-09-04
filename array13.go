package main

import "fmt"

func main() {
	var array = [7]int{1, 2, 3, 4, 5, 6, 7}
	var numero int
	novoArray := make([]int, 7)

	fmt.Println(array)
	fmt.Println("Digite o numero que voce quer que fique em primeiro e em ultimo nesta lista")
	fmt.Scanln(&numero)

	for i := 0; i < len(array); i++ {
		array[0] = numero
		array[4] = numero
		novoArray = append(novoArray, array[i])

	}
}