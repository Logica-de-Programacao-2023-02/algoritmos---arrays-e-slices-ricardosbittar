package main

import "fmt"

func main() {
	var produto float64
	array := [5]float64{1.2, 2.6, 3.5, 4.4, 5.1}

	for i := 0; i < len(array); i++ {

		produto *= array[i]

	}
	fmt.Println(produto)
}
