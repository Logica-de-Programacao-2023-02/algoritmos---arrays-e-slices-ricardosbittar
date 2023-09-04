package main

import "fmt"

func main() {

	var linhas [3]int
	var colunas [2]int
	var numLinhas int
	var numColunas int
	var nLinhas []int
	var nColunas []int

	for i := 0; i < len(linhas); i++ {
		fmt.Println("Digite um número")
		fmt.Scanln(&numLinhas)

		nLinhas = append(nLinhas, numLinhas)
	}
	for i := 0; i < len(colunas); i++ {
		fmt.Println("Digite um número")
		fmt.Scanln(&numColunas)

		nColunas = append(nColunas, numColunas)

	}

	var matriz [][]int
	matriz = append(matriz, nLinhas, nColunas)
	fmt.Print(matriz)

}
