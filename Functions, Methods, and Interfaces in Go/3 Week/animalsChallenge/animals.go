package main

import "fmt"

var colorRed = "\033[31m"
var colorReset = "\033[0m"
var colorBlue = "\033[34m"
var colorCyan = "\033[36m"

/*Animal class */
type Animal struct {
	food       string
	locomotion string
	noise      string
}

/*Eat is an Animal method */
func (animal *Animal) Eat() {
	fmt.Println(string(colorRed), animal.food)
}

/*Move is an Animal method */
func (animal *Animal) Move() {
	fmt.Println(string(colorBlue), animal.locomotion)
}

/*Speak is an Animal method*/
func (animal *Animal) Speak() {
	fmt.Println(string(colorCyan), animal.noise)

}

func animalAction(animal Animal) {
	animalAction := ""
	fmt.Println(string(colorReset), "Which info do u want?")
	_, _ = fmt.Scan(&animalAction)
	switch animalAction {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}

}

func main() {
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	condition := true
	var action string
	for ok := true; ok; ok = condition {
		fmt.Println(string(colorReset), "Type an animal:")
		fmt.Scan(&action)
		switch action {
		case "cow":
			animalAction(cow)
		case "bird":
			animalAction(bird)
		case "snake":
			animalAction(snake)
		}

	}
}
