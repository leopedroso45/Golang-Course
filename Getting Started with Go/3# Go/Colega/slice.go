package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tmp []int
	s := make([]int, 0, 3)
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter integer or X")
		if line, err := in.ReadString('\n'); err == nil {
			line = strings.TrimRight(line, "\r\n")
			if line == "X" {
				return
			}
			if i, err := strconv.Atoi(line); err == nil {
				if len(s) == 0 {
					s = append(s, i)
				} else {
					for index, num := range s {
						if num > i {
							tmp = make([]int, len(s)+1)
							copy(tmp, s[:index])
							tmp = append(tmp[:index], i)
							s = append(tmp, s[index:]...)
							break
						}
						if index == len(s)-1 {
							s = append(s, i)
						}
					}
				}
				fmt.Println(s)
			}
		}
	}
}
