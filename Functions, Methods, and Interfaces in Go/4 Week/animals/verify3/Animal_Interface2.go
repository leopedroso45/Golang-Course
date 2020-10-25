package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var atype, name1, strLower1, strLower2, strLower3 string
var newAnimal AnimalNew

type AnimalNew interface {
	Eat() string
	Move() string
	Speak() string
}
type Cow struct {
}

func (c Cow) Name() string {
	return ("The Name of the cow is :" + name1)
}

func (c Cow) Eat() string {
	return ("The cow eats Grass")
}
func (c Cow) Move() string {
	return ("The Cow Walks")
}
func (c Cow) Speak() string {
	return ("The cow Moos")
}

type Bird struct {
}

func (b Bird) Name() string {
	return ("The Name of the Bird is :" + name1)
}

func (b Bird) Eat() string {
	return ("The Bird eats worms")
}
func (b Bird) Speak() string {
	return ("The Bird Chirps")
}
func (b Bird) Move() string {
	return ("The Bird Flies")
}

type Snake struct {
}

func (s Snake) Name() string {
	return ("The Name of the Snake is :" + name1)
}
func (s Snake) Eat() string {
	return ("The Snake eats mice")
}
func (s Snake) Speak() string {
	return ("The Snake Hisses")
}
func (s Snake) Move() string {
	return ("The Snake Slithers")
}

func stripSpaces(str_in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str_in)
}
func query() (string, error) {

	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(" This Program takes three inputs -query,  name of animal, and action type of the animal {Eat, Move, Speak }")
	fmt.Println(" To Exit the Program - please press Cntrl + C and hit enter")
	fmt.Println(" Please enter the values from the above list you wish to see seperated by a space: >")
	reader2 := bufio.NewReader(os.Stdin)
	strInput2, _ := reader2.ReadString('\n')
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------")

	var eat1 string = "eat"
	var move1 string = "move"
	var speak1 string = "speak"
	name1 = ""
	strIn2 := strings.Split(strInput2, " ")
	strLower1 = strings.ToLower(strIn2[0])
	strLower2 = strings.ToLower(strIn2[1])
	strLower3 = strings.ToLower(strIn2[2])

	name1 = stripSpaces(strLower2)
	acttype := stripSpaces(strLower3)

	switch {
	case acttype == eat1:
		fmt.Println(newAnimal.Eat())
	case acttype == move1:
		fmt.Println(newAnimal.Move())
	case acttype == speak1:
		fmt.Println(newAnimal.Speak())
	}

	return name1, errors.New("Please check the Instructions Clearly and check the Input")

}
func resetVars() {
	atype = ""
	name1 = ""
	strLower1 = ""
	strLower2 = ""
	strLower3 = ""
	newAnimal = nil
}

func newAnimalInput() (string, error) {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(" Enter a command  newanimal, the name and type of animal :(cow, bird , snake) seperated by a space ")
	fmt.Println(" To Exit the Program - please press Cntrl + C and hit enter")
	fmt.Println(" Please enter the values from the above list you wish to see seperated by a space: >")
	reader := bufio.NewReader(os.Stdin)
	strInput, _ := reader.ReadString('\n')
	strInput = strings.ToLower(strInput)
	strIn1 := strings.Split(strInput, " ")
	strLower1 := strings.ToLower(strIn1[0])
	strLower2 := strings.ToLower(strIn1[1])
	strLower3 := strings.ToLower(strIn1[2])

	if "newanimal" == strLower1 {

		var cowtype string = "cow"
		var birdtype string = "bird"
		var snaketype string = "snake"

		name1 = stripSpaces(strLower2)
		atype = stripSpaces(strLower3)

		switch {
		case atype == cowtype:
			var cow1 Cow
			newAnimal = cow1

		case atype == birdtype:
			var bird1 Bird
			newAnimal = bird1

		case atype == snaketype:
			var snk1 Snake
			newAnimal = snk1
		}
	} else if "query" == strLower1 {
		newAnimalInput()
		fmt.Println("Created It !!!")
		query()
		resetVars()
	}
	return name1, errors.New("Please Read the Instructions Clearly and check the Input")
}

func main() {
	for {
		newAnimalInput()
		fmt.Println("Created It !!!")
		if atype != "" {
			query()
			resetVars()
		} else {
			fmt.Println(errors.New("Please Read the Instructions Clearly and check the Input"))
		}
	}
} // end
