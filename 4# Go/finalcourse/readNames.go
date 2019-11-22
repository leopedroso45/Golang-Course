package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	type Person struct {
		fname string
		lname string
	}
	var PersonSlice = make([]Person, 0, 3)

	fmt.Printf("Hello my friend, what's the file name?")
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		personString := strings.Fields(scanner.Text())
		fname := personString[0]
		lname := personString[1]

		p1 := Person{fname: fname, lname: lname}
		PersonSlice = append(PersonSlice, p1)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(PersonSlice)
}
