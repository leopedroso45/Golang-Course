package main

import "fmt"

func main() {
	fmt.Println(getMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 100))
	vslice := []int{1, 3, 6, 4}
	fmt.Println(getMax(vslice...))
}

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}
