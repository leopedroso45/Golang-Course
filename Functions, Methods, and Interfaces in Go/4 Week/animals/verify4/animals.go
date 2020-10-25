package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Speak()
	Eat()
	Move()
}

type animal_characteristics struct {
	food       string
	locomotion string
	noise      string
}

type Cow struct {
	name            string
	characteristics animal_characteristics
}

type Bird struct {
	name            string
	characteristics animal_characteristics
}

type Snake struct {
	name            string
	characteristics animal_characteristics
}

func (a Cow) Eat() {
	fmt.Println(a.characteristics.food)
}

func (a Cow) Move() {
	fmt.Println(a.characteristics.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.characteristics.noise)
}

func (a Bird) Eat() {
	fmt.Println(a.characteristics.food)
}

func (a Bird) Move() {
	fmt.Println(a.characteristics.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.characteristics.noise)
}

func (a Snake) Eat() {
	fmt.Println(a.characteristics.food)
}

func (a Snake) Move() {
	fmt.Println(a.characteristics.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.characteristics.noise)
}

func read_request() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter 'x' to end loop or  newanimal command or query command >")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}

func parse_request(text string) (string, string, string) {
	h := strings.Split(text, " ")
	command := h[0]
	name := h[1]
	type_act := h[2]
	return command, name, type_act

}

func query_action(a Animal, act string) {
	switch act {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("wrong action:" + act)
	}
}

func create_animal(a_name, type_act string) Animal {
	var a Animal
	switch type_act {
	case "cow":
		cow_c := animal_characteristics{food: "grass", locomotion: "walk", noise: "moo"}
		a = Cow{name: a_name, characteristics: cow_c}
	case "bird":
		bird_c := animal_characteristics{food: "worms", locomotion: "fly", noise: "peep"}
		a = Bird{name: a_name, characteristics: bird_c}
	case "snake":
		snake_c := animal_characteristics{food: "mice", locomotion: "slither", noise: "hsss"}
		a = Snake{name: a_name, characteristics: snake_c}

	default:
		fmt.Println("wrong type:" + type_act)
	}

	return a

}

func main() {
	var animals_m = make(map[string]Animal)

Loop:
	for {

		request := read_request()

		switch request {
		case "X":
			break Loop
		case "x":
			break Loop
		default:

			command, name, type_act := parse_request(request)
			//fmt.Println(command, name, type_act)

			switch command {
			case "newanimal":

				animals_m[name] = create_animal(name, type_act)
				//fmt.Println(animals_m)
				fmt.Println("Created it!")

			case "query":
				if val, ok := animals_m[name]; ok {
					query_action(val, type_act)
				} else {
					fmt.Println(name + " not found")
				}
			default:
				fmt.Println("Wrong input:" + request)
			}

		}
	}

}
