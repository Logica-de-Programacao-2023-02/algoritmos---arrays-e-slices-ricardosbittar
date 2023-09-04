package main

func main() {

	array := [5]int{1, 2, 3, 4, 5}
	var slice []int
	for i := 0; i < len(array); i++ {
		if i%3 == 0 {
			slice = append(slice, array[i])
		} else {
			continue
		}
	}
}
