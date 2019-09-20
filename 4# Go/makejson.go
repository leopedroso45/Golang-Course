package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello my friend, could you inform your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("I need to know your address too...")
	var address string
	fmt.Scanln(&address)

	peopleMap := map[string]string{
		"name":    name,
		"address": address,
	}
	barr, err := json.Marshal(peopleMap)
	f, err := os.Create("test.json")
	if err == nil {
		_, err := f.Write(barr)
		if err == nil {
			fmt.Println("Done!")
			for key, val := range peopleMap {
				fmt.Println("'" + key + "'" + ":" + "'" + val + "'")
			}
		}
	}
}
