package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("idk.txt")
	if err == nil {
		barr := []byte{1, 0, 1}
		_, err := f.Write(barr)
		_, err = f.WriteString("Ai meu Deus!")
		if err == nil {
			fmt.Printf("foi foi foi")
		}
	}
}
