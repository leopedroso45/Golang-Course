package main

import "fmt"

func foo(x int, y int) int {
	return x * y
}

func foo2(x, y int) int {
	return x / y
}

func foo3(x int) (int, int) {
	return x, x + 1
} 

func main() {
	fmt.Println(foo(2, 3))
	fmt.Println(foo2(5, 5))
	fmt.Println(foo3(5))

}
