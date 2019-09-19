package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	dat, err := ioutil.ReadFile("test.txt")
	if err == nil {
		fmt.Println(dat)
	}
	err = ioutil.WriteFile("outfile.txt", dat, 0777)

}
