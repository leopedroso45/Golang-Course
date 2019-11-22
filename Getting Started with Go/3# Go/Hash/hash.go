package main

import (
	"fmt"
)

func main() {
	//var idMap map[string]int
	idMap := make(map[string]int)
	fmt.Println(idMap)

	idMap2 := map[string]string{
		"Name":    "Leonardo",
		"Country": "Brasil"}
	//fmt.Println(idMap2)
	fmt.Println(idMap2["Name"])
	//Increase
	idMap2["Family's name"] = "Severo"
	fmt.Println(idMap2)
	//delete
	delete(idMap2, "Family's name")
	fmt.Println(idMap2)
	//two-value assignment tests for existence os the key
	id, p := idMap2["Name"]
	fmt.Println(id)
	fmt.Println(p)
	fmt.Println(len(idMap2))
	//Iterating through a Map
	//using a for loop with the range keyword
	//two value assignment with range
	for key, val := range idMap2 {
		fmt.Println(key, val)
	}
}
