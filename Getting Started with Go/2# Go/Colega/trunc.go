package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {

	var i int = 0

	for i = 0; i < 2; i++ {
		fmt.Printf("Truncated number: %d\n", truncateFloatInput(readInput()))
	}

}

func readInput() string {

	var input string

	fmt.Printf("Enter a floating point number: ")

	_, err := fmt.Scan(&input)

	if (err != nil) {
		fmt.Printf("Error reading input: %s", err)
		os.Exit(1)
	}

	return input

}

func truncateFloatInput(input string) uint64 {

	number, err := strconv.ParseFloat(input, 64)

	if (err != nil) {
		fmt.Printf("Error converting input to float: %s", err)
		os.Exit(1)
	}

	return uint64(number)

}
