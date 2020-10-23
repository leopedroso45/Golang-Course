package main

import "fmt"

type speaker interface {
	speak()
}

type dog struct {
	name string
}

func (d dog) speak() {
	fmt.Println(d.name)
}

func main() {
	var s1 speaker
	d1 := dog{
		name: "brian",
	}
	s1 = d1
	s1.speak()

}
