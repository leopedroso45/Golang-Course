package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var String
	var stringValue string
	var originalValue string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("please, type some text:")
	//user input
	stringValue, _ = reader.ReadString('\n')
	stringValue = strings.TrimSuffix(stringValue, "\n")
	originalValue = stringValue
	//remove spaces
	stringValue = strings.ReplaceAll(stringValue, " ", "")
	//set string to lower case
	stringValue = strings.ToLower(stringValue)
	//if contains substring "i", "a" and "n"
	if strings.Contains(stringValue, "i") && strings.Contains(stringValue, "a") && strings.Contains(stringValue, "n") {
		// if "i" is a prefix and "n" a suffix
		// ------- here
		firstChar := strings.HasPrefix(stringValue, "i")
		lastChar := strings.HasSuffix(stringValue, "n")
		switch {
		//if is a prefix and suffix so Found
		case firstChar && lastChar:
			fmt.Printf("Found in %s", originalValue)
		//if isnt a prefix and suffix so Not Found
		default:
			fmt.Printf("Not Found in %s", originalValue)
		}
		//if there's no "i", "a" and "n"
	} else {
		fmt.Printf("Not Found in %s", originalValue)
	}
}
