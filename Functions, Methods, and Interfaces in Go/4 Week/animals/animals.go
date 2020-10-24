package main

import "fmt"

type animal interface {
	speak()
	eat()
	move()
}

type cow struct {
	name       string
	locomotion string
	noise      string
	eaten      string
}

func (c *cow) initit() {
	c.locomotion = "walk"
	c.noise = "moo"
	c.eaten = "grass"
}

func (c *cow) speak() {
	fmt.Println(c.noise)
}
func (c *cow) eat() {
	fmt.Println(c.eaten)
}
func (c *cow) move() {
	fmt.Println(c.locomotion)
}

type bird struct {
	name       string
	locomotion string
	noise      string
	eaten      string
}

func (b *bird) initit() {
	b.locomotion = "fly"
	b.noise = "peep"
	b.eaten = "worms"
}

func (b *bird) speak() {
	fmt.Println(b.noise)
}
func (b *bird) eat() {
	fmt.Println(b.eaten)
}
func (b *bird) move() {
	fmt.Println(b.locomotion)
}

type snake struct {
	name       string
	locomotion string
	noise      string
	eaten      string
}

func (s *snake) initit() {
	s.locomotion = "slither"
	s.noise = "hsss"
	s.eaten = "mice"

}

func (s *snake) speak() {
	fmt.Println(s.noise)
}
func (s *snake) eat() {
	fmt.Println(s.eaten)
}
func (s *snake) move() {
	fmt.Println(s.locomotion)
}

var (
	cows   []cow
	birds  []bird
	snakes []snake
)

func main() {
	condition := true
	var action string
	for ok := true; ok; ok = condition {
		fmt.Println("Type a request, it can be a 'query' or 'newanimal':")
		fmt.Scan(&action)
		switch action {
		case "query":
			query()
		case "newanimal":
			newanimal()
		}

	}
}

func newanimal() {
	var name string
	fmt.Println("Inform the animal name")
	fmt.Scan(&name)
	fmt.Println("Inform the animal type")
	var typeAnimal string
	fmt.Scan(&typeAnimal)
	switch typeAnimal {
	case "cow":
		cowish := cow{
			name: name,
		}
		cowish.initit()
		cows = append(cows, cowish)
		fmt.Println("Created it!")
		return
	case "bird":
		birdish := bird{
			name: name,
		}
		birdish.initit()
		birds = append(birds, birdish)
		fmt.Println("Created it!")
	case "snake":
		snakeish := snake{
			name: name,
		}
		snakeish.initit()
		snakes = append(snakes, snakeish)
		fmt.Println("Created it!")
	default:
		fmt.Println("I can't understand you, try again.")
		return
	}
}

func query() {
	var name string
	fmt.Println("Inform the animal name")
	fmt.Scan(&name)
	fmt.Println("Inform the animal type")
	var typeAnimal string
	fmt.Scan(&typeAnimal)
	fmt.Println("Inform the animal action")
	var action string
	fmt.Scan(&action)
	switch typeAnimal {
	case "cow":
		for _, cow := range cows {
			if cow.name == name {
				if action == "eat" {
					cow.eat()
					return
				} else if action == "move" {
					cow.move()
					return
				} else if action == "speak" {
					cow.speak()
					return
				}
			} else {
				fmt.Println("Can't find it")
			}
		}
	case "bird":
		for _, bird := range birds {
			if bird.name == name {
				if action == "eat" {
					bird.eat()
					return
				} else if action == "move" {
					bird.move()
					return
				} else if action == "speak" {
					bird.speak()
					return
				}
			} else {
				fmt.Println("Can't find it")
			}
		}
	case "snake":
		for _, snake := range snakes {
			if snake.name == name {
				if action == "eat" {
					snake.eat()
					return
				} else if action == "move" {
					snake.move()
					return
				} else if action == "speak" {
					snake.speak()
					return
				}
			} else {
				fmt.Println("Can't find it")
			}
		}
	default:
		fmt.Println("Can't find it.")
		return
	}
}
