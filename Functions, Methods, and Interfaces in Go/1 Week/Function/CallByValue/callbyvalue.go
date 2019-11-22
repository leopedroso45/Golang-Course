package main

import "fmt"

func foo(x int) {
	x = x + 1
}

func main() {
	x := 1
	foo(x)
	fmt.Println(x)

}
