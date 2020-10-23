package main

import (
	"fmt"
	"math"
)

var (
	acceleration        float64
	initialVelocity     float64
	initialDisplacement float64
	time                float64
)

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*(math.Pow(t, 2)*0.5) + (v0 * t) + (s0))
	}
	return fn
}

func main() {

	fmt.Printf("Type a value for acceleration: ")
	_, _ = fmt.Scan(&acceleration)
	fmt.Printf("Type a value for initial velocity:")
	_, _ = fmt.Scan(&initialVelocity)
	fmt.Printf("Type a value for initial displacement:")
	_, _ = fmt.Scan(&initialDisplacement)
	fmt.Printf("Type a value for time:")
	_, _ = fmt.Scan(&time)

	fn := genDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(fn(time))

}
