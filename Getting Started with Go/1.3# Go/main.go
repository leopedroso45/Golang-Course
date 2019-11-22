package main

import "fmt"

func main() {
	var appleNum int

	fmt.Printf("Number of apples?")

	fmt.Scan(&appleNum)

	fmt.Printf("%d", appleNum)
}
