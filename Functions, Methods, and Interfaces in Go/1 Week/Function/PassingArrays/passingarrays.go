package main

import "fmt"

func foo(x [3]int) int {
	return x[0]
}

//Wrong way using an array
func foo2(x *[3]int) int {
	(*x)[0] = 7
	return (*x)[0]
}

//Right way using a slice
func foo3(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := []int{1, 2, 3}
	foo2(&b)
	foo3(c)
	fmt.Println(foo(a))
	fmt.Println(b[0])
	fmt.Println(c)
}
