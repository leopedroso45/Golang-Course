package main

import (
	"fmt"
)

func main() {

	var x [5]string
	tempInt := 0
	y := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(x); i++ {
		x[i] = "Test"
		y[i] = tempInt + 1
		tempInt++
	}
	//dont need a method
	fmt.Println(x) //Super beauti
	fmt.Println(y) //Super beauti
}
