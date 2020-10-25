package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorCyan   = "\033[36m"
	colorReset  = "\033[0m"
)

/*Animal interface contain Eat, Move and Speak methods. */
type Animal interface {
	Eat()
	Move()
	Speak()
	InitIt(string) Animal
	GetName() string
}

/*Cow animal type that implements Animal */
type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

/*Eat animal method */
func (a *Cow) Eat() {
	fmt.Println(string(colorGreen), a.name+": "+a.food, string(colorReset))
}

/*Move animal method */
func (a *Cow) Move() {
	fmt.Println(string(colorGreen), a.name+": "+a.locomotion, string(colorReset))
}

/*Speak animal method */
func (a *Cow) Speak() {
	fmt.Println(string(colorGreen), a.name+": "+a.noise, string(colorReset))
}

/*InitIt animal method */
func (a *Cow) InitIt(name string) Animal {
	a.name = name
	a.food = "grass"
	a.locomotion = "walk"
	a.noise = "moo"
	fmt.Println(string(colorCyan), "Created It!", string(colorReset))
	return a
}

/*GetName animal method */
func (a *Cow) GetName() string {
	return a.name
}

/*Bird animal type that implements Animal */
type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

/*Eat animal method */
func (a *Bird) Eat() {
	fmt.Println(string(colorYellow), a.name+": "+a.food, string(colorReset))
}

/*Move animal method */
func (a *Bird) Move() {
	fmt.Println(string(colorYellow), a.name+": "+a.locomotion, string(colorReset))
}

/*Speak animal method */
func (a *Bird) Speak() {
	fmt.Println(string(colorYellow), a.name+": "+a.noise, string(colorReset))
}

/*InitIt animal method */
func (a *Bird) InitIt(name string) Animal {
	a.name = name
	a.food = "worms"
	a.locomotion = "fly"
	a.noise = "peep"
	fmt.Println(string(colorCyan), "Created It!", string(colorReset))
	return a
}

/*GetName animal method */
func (a *Bird) GetName() string {
	return a.name
}

/*Snake animal type that implements Animal */
type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

/*Eat animal method */
func (a *Snake) Eat() {
	fmt.Println(string(colorRed), a.name+": "+a.food, string(colorReset))
}

/*Move animal method */
func (a *Snake) Move() {
	fmt.Println(string(colorRed), a.name+": "+a.locomotion, string(colorReset))
}

/*Speak animal method */
func (a *Snake) Speak() {
	fmt.Println(string(colorRed), a.name+": "+a.noise, string(colorReset))
}

/*InitIt animal method */
func (a *Snake) InitIt(name string) Animal {
	a.name = name
	a.food = "mice"
	a.locomotion = "slither"
	a.noise = "hsss"
	fmt.Println(string(colorCyan), "Created It!", string(colorReset))
	return a
}

/*GetName animal method */
func (a *Snake) GetName() string {
	return a.name
}

var animals []Animal

func main() {
	keepAlive := true
	var action string
	for ok := true; ok; ok = keepAlive {
		fmt.Println("Use 'newanimal' or 'query':")
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			action = scanner.Text()
			break
		}
		commands := strings.Fields(action)

		if commands[0] == "query" {
			for _, a := range animals {
				n := a.GetName()
				np := commands[1]
				cmd := commands[2]
				if n == np {
					switch cmd {
					case "eat":
						a.Eat()
					case "move":
						a.Move()
					case "speak":
						a.Speak()
					default:
						fmt.Println("Can't undestand. Try again.")
						return
					}
				}
			}
		} else if commands[0] == "newanimal" {
			animalType := commands[2]
			switch animalType {
			case "cow":
				var c Cow
				var a Animal
				a = c.InitIt(commands[1])
				animals = append(animals, a)
			case "bird":
				var c Bird
				var a Animal
				a = c.InitIt(commands[1])
				animals = append(animals, a)
			case "snake":
				var c Snake
				var a Animal
				a = c.InitIt(commands[1])
				animals = append(animals, a)
			default:
				fmt.Println("Can't understand. Try again.")
				return
			}

		} else {
			fmt.Println("I can't understand. Try again.")
		}
	}
}
