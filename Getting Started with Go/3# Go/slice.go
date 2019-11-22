package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var stringValue string
var slice = make([]int64, 0, 3)

func main() {

	for true {
		fmt.Printf("Please enter an integer value to continue or 'x' to quit:")
		fmt.Scan(&stringValue)
		//Make sure stringValue can be an integer and convert it
		if _, err := strconv.Atoi(stringValue); err == nil {
			intValue := convertToInt(stringValue)
			//Append the integer value
			slice = append(slice, intValue)
			//Sort the slice using the sort method
			sort.Slice(slice, func(i, j int) bool {
				return slice[i] < slice[j]
			})
			fmt.Println(slice)
		} else {
			if stringValue == "X" {
				fmt.Println("See you later... Bye!")
				os.Exit(0)
			} else {
				fmt.Println("Please enter a valid value")
			}
		}
	}
}

func convertToInt(stringRcv string) int64 {
	//Parse to Integer
	intGer, err := strconv.ParseInt(stringRcv, 10, 64)
	if err == nil {
		return intGer
	}
	return intGer
}
