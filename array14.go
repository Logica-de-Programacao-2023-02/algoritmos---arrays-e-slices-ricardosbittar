package main

import "fmt"

func main() {
	slice := make([]int, 8)
	var i1, i2 int
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Informe o valor para a posição %d: ", i)
		fmt.Scanln(&slice[i])
	}
	fmt.Print("Informe o primeiro índice a ser trocado: ")
	fmt.Scanln(&i1)
	fmt.Print("Informe o segundo índice a ser trocado: ")
	fmt.Scanln(&i2)

	slice[i1], slice[i2] = slice[i2], slice[i1]
	fmt.Println(slice)
}
