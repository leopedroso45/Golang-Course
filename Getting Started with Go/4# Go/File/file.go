package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	dat, err := ioutil.ReadFile("test.txt")
	if err == nil {
		s := string(dat)
		fmt.Println(s)
	}
	err = ioutil.WriteFile("outfile.txt", dat, 0777)

}
