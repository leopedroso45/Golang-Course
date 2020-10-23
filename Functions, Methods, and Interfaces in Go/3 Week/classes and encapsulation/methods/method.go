package main

import "fmt"

type human struct {
	name      string
	isWalking bool
}

//we could create getters and setters like this
func (h *human) walk() {
	fmt.Println(h.name + " is walking.")
	h.isWalking = true
}

func (h *human) stopWalk() {
	fmt.Println(h.name + " is looking at the corner.")
	h.isWalking = false
}

func main() {
	var humano human
	fmt.Scan(&humano.name)
	humano.isWalking = false
	condition := true
	var action string
	for ok := true; ok; ok = condition {
		fmt.Scan(&action)
		switch action {
		case "walk":
			humano.walk()
		case "stop":
			humano.stopWalk()
		}

	}

}
