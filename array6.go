package main

import "fmt"

func main() {
	var numero int
	var array = [10]int{1, 4, 8, 6, 8, 12, 5, 2, 9, 10}

	fmt.Println("Digite um número")
	fmt.Scanln(&numero)
	for i := 0; i < len(array); i++ {
		if i == numero {
			fmt.Println("O numero que voce digitou esta na lista")
			break
		} else {
			fmt.Println("O numero que voce digitou nao esta na lista")
			continue
		}
	}
}
