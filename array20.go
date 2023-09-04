package main

import "fmt"

func main() {

	var array = []int{1, 2, 3, 4, 5}
	crescente := true

	for i := 0; i < len(array); i++ {
		if array[i] > array[i+1] {

			crescente = false
		}

	}
	if crescente {
		fmt.Println("a lista esta em ordem crescente")
	} else {
		fmt.Println("A lista nao esta em ordem crescente")
	}
}