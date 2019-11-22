package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]string)
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name:")
	if line, err := in.ReadString('\n'); err == nil {
		m["name"] = strings.Replace(line, "\r\n", "", -1)
	}
	fmt.Println("Enter address:")
	if line, err := in.ReadString('\n'); err == nil {
		m["address"] = strings.Replace(line, "\r\n", "", -1)
	}
	if result, err := json.Marshal(m); err == nil {
		fmt.Println(string(result))
	}
}
