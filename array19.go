package main

import "fmt"

func main() {
	var numero int
	var num1 int
	var num2 int
	var soma int

	fmt.Println("Informe um numero que sera o tamanho das duas listas ")
	fmt.Scanln(&numero)
	for i:= 0; i<numero{
	fmt.Println("Informe os numeros que estarao presentes na primeira lista")
	fmt.Scanln(&num1)
	fmt.Println("Informe os numeros que estarao presentes na segunda lista")
	fmt.Scanln(&num2)

	 array1 := []int{num1}
	 array2:= []int{num2}
	 var array3 []int
	 array3= append(array3, array1, array2)


}
