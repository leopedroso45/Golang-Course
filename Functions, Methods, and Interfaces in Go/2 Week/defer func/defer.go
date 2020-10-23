package main

import "fmt"

func main() {
	defer fmt.Println("Bye!")
	fmt.Println("Hello!!")

	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println(i + 1)
}
