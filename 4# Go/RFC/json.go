package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var p2 Person
	err = json.Unmarshal(barr, &p2)
	fmt.Printf(p2.name, p2.addr, p2.phone)

	err = ioutil.WriteFile("outfile.json", barr, 0777)

}
