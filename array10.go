package main

import "fmt"

func main() {
	slice := make([]int, 10)
	var maior int
	slice = append(slice, 8, 4, 6, 3, 12, 9, 7, 44, 5, 20)
	for i := 0; i < len(slice); i++ {
		if i > maior {
			maior = i
		}
		fmt.Println(maior)
	}
}
