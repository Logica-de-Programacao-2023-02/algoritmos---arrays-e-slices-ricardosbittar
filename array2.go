package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	novoSlice := append(slice[:2], slice[3:]...)
	fmt.Println(novoSlice)
}
