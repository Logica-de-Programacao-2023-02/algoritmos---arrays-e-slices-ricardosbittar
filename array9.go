package main

import "fmt"

func main() {
	array := make([]float64, 6)
	var valor float64

	fmt.Println("Informe um valor")
	fmt.Scanln(&valor)
	for i := 0; i < len(array); i++ {
		array[i] = valor

	}
	novoSlice := make([]float64, 6)
	novoSlice = append(novoSlice, valor)
}
