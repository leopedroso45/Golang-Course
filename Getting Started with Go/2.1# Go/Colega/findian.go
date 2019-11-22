package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please enter a string: ")

	// In the first version of program I used fmt.Scan
	// but on my Mac if occurred an error with whitespaced values
	// like "I d skd a efju N", so I Googled how to read raw stdin via bufio package
	cliReader := bufio.NewReader(os.Stdin)
	enteredValue, _ := cliReader.ReadString('\n')

	enteredValue = strings.Trim(enteredValue, "\n")
	enteredValue = strings.ToLower(enteredValue)

	if strings.HasPrefix(enteredValue, "i") &&
		strings.Contains(enteredValue, "a") &&
		strings.HasSuffix(enteredValue, "n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
