package main

import "fmt"

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func x2(number int) int {
	return number * 2
}

func x3(number int) int {
	return number * 3
}

func main() {
	soma1 := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println(soma1)
	fmt.Println(applyIt(x2, 2))
	fmt.Println(applyIt(x3, 2))

}
