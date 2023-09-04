package main

import "fmt"

func main() {
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int
	fmt.Print("Informe o índice da linha (0 ou 1): ")
	fmt.Scanln(&linha)
	fmt.Print("Informe o índice da coluna (0, 1 ou 2): ")
	fmt.Scanln(&coluna)

	if linha >= 0 && linha < 2 && coluna >= 0 && coluna < 3 {

		valor := matriz[linha][coluna]
		fmt.Printf("Valor na posição (%d, %d): %d\n", linha, coluna, valor)
	} else {
		fmt.Println("Índices inválidos. Certifique-se de que estão dentro do intervalo permitido.")
	}
}
