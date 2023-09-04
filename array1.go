package main

import "fmt"

func main() {
	soma := 0
	array := [3]int{1, 2, 3}

	for i := 0; i < len(array); i++ {
		soma += array[i]

	}
	fmt.Println(soma)

}
