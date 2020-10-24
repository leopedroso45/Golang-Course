package main

import (
	"fmt"
)

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		var command, name, attr string
		fmt.Scan(&command, &name, &attr)
		if command == "query" {
			animal, ok := animals[name]
			if ok {
				switch attr {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Entered wrong information type. Use: eat, move, speak.")
				}
			} else {
				fmt.Println("% not found", animal)
			}
		} else if command == "newanimal" {
			animals[name] = animals[attr]
			fmt.Println("created!")
		}
	}
}

type Animal struct {
	food, locomotion, sound string
}

type AnimalAttrs interface {
	Eat()
	Move()
	Speak()
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}
