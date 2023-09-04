package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe um numero inteiro")
	fmt.Scanln(&n)
	var slice []int

	n = len(slice)

	for i := 0; i < n; i++ {
		if i%1 == 0 && i%i == 0 {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)

}
