package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type namestruct struct {
	firstname  string
	secondname string
}

func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	var fileName string
	names := []*namestruct{}
	fmt.Println("Input file name")
	fileName = scanInput()

	x, _ := scanLines(fileName)
	for _, d := range x {
		a := strings.Split(d, " ")
		names = append(names, &namestruct{firstname: a[0], secondname: a[1]})
	}
	for _, d := range names {
		fmt.Printf("%v", d)
	}
}
