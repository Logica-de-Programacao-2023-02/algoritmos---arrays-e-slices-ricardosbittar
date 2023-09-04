package main

import "fmt"

func main() {

	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	soma := array[0] + array[2] + array[4] + array[6] + array[8]
	fmt.Println(soma)
}
