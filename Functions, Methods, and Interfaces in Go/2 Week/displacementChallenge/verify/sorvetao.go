package main

import "fmt"

func GenDisplaceFn(a, v0, d0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a/2)*t*t + v0*t + d0
	}
	return fn
}

func main() {
	var a, v0, d0, t float64
	fmt.Println("a=")
	fmt.Scanln(&a)
	fmt.Println("v0=")
	fmt.Scanln(&v0)
	fmt.Println("d0=")
	fmt.Scanln(&d0)
	fmt.Println("t=")
	fmt.Scanln(&t)

	fn := GenDisplaceFn(a, v0, d0)
	fmt.Println(fn(t))
}
