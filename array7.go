package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	var slice []int
	var numero int
	fmt.Println("Ïnforme um numero: ")
	fmt.Scanln(&numero)
	for i := 0; i < len(array); i++ {
		if numero == i {
			continue
		} else {
			slice = append(slice, array[i], numero)
		}

	}
}
