package main

import "fmt"

func main() {
	slice := make([]string, 8)
	var valor string
	fmt.Println("Informe um valor ")
	fmt.Scanln(&valor)
	slice[0] = "joao"
	slice[1] = "Maria"
	slice[2] = "Jose"
	slice[3] = "Pedro"
	slice[4] = "Ricardo"
	slice[5] = "Gabriel"
	slice[6] = "Joaquim"
	slice[7] = "Arthur"

	novoSlice := removerOcorrencias(slice, valor)

	fmt.Println("Slice original:", slice)
	fmt.Println("Slice após a remoção das ocorrências de", valor, ":", novoSlice)
}

func removerOcorrencias(slice []string, valor string) []string {
	resultado := []string{}
	for _, elemento := range slice {
		if elemento != valor {
			resultado = append(resultado, elemento)
		}
	}
	return resultado
}
