package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Animal struct{
	eat string
	move string
	speak string
}

func (animal *Animal) Eat(){
	fmt.Printf("Eat:%s\n",animal.eat)
}

func (animal *Animal) Move(){
	fmt.Printf("Move:%s\n",animal.move)
}
func (animal *Animal) Speak(){
	fmt.Printf("Speak:%s\n",animal.speak)
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	cow := Animal{"grass","walk","moo"}
	bird := Animal{"worms","fly","peep"}
	snake := Animal{"mice","slither","hsss"}
	var input []string
	fmt.Println("Provide two space separated(first name of animal,second information needed) argument after >(Enter X to stop)")
	for{
		fmt.Print(">")
		input1,_ := reader.ReadString('\n')
		input1 = strings.TrimRight(input1, "\r\n")
		input1 = strings.ToLower(input1)
		input = strings.Split(input1, " ") 
		if (input[0] == "x"){
			break
		}else{
			if (input[0] == "cow"){
				fmt.Println("Name:Cow")
				if (input[1] == "eat"){
					cow.Eat()
				}else if (input[1] == "move"){
					cow.Move()
				}else if (input[1] == "speak"){
					cow.Speak()
				}else{
					fmt.Println("Incorrect Input")
				}
			}else if (input[0] == "bird"){
				fmt.Println("Name:Bird")
				if (input[1] == "eat"){
					bird.Eat()
				}else if (input[1] == "move"){
					bird.Move()
				}else if (input[1] == "speak"){
					bird.Speak()
				}else{
					fmt.Println("Incorrect Input")
				}
			}else if (input[0] == "snake"){
				fmt.Println("Name:Snake")
				if (input[1] == "eat"){
					snake.Eat()
				}else if (input[1] == "move"){
					snake.Move()
				}else if (input[1] == "speak"){
					snake.Speak()
				}else{
					fmt.Println("Incorrect Input")
				}
			}else{
				fmt.Println("Incorrect Input")
			}
		}
	}
}


