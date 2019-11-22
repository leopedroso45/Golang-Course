package main

import (
	"fmt"
)

func main() {

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	//len value number and cap the total cap of an array
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(arr), cap(arr))
	//I can change the slice value
	s1[1] = "asb"
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	sli := []int{1, 2, 3}
	fmt.Println(sli)
	s := make([]string, 3)
	//	s := make([]string, 3, 15)
	fmt.Println("emp:", s)
	sli3 := make([]int, 0, 3)
	//sli3 = append(sli3, 100)
	fmt.Println(len(sli3), cap(sli3))

}
