package main

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int
	for i := 0; i < len(array); i++ {
		if i%2 == 0 {
			slice = append(slice, array[i])
		} else {
			continue
		}
	}
}
