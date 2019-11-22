package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Person struct {
		name  string
		addr  string
		phone string
	}

	p1 := Person{name: "Leonardo", addr: "A st.", phone: "123"}
	barr, err := json.Marshal(p1)
	fmt.Println(barr, err)
}
