package main

import (
	"fmt"
)

var swapped = true

func main() {
	numbers := GetNumbers()
	BubbleSort(numbers)
}

/*BubbleSort function*/
func BubbleSort(input []int) {
	n := 10
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				Swap(input, i)
			}
		}
	}
	fmt.Println(input)
}

func Swap(input []int, i int) {
	input[i], input[i-1] = input[i-1], input[i]
	swapped = true
}

/*GetNumbers get ten numbers from the user and it returns a slice*/
func GetNumbers() []int {
	var intReceived int
	var numbersReceived []int

	for i := 0; i < 10; i++ {
		fmt.Printf("Type an integer please:")
		valor, _ := fmt.Scan(&intReceived)
		if valor != 0 {
			numbersReceived = append(numbersReceived, intReceived)
		} else {
			i--
		}
	}
	return numbersReceived
}
