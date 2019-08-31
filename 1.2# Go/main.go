package main

import "fmt";

func main() {
	res1 := Somar(1, 2)
	res2 := Subtrair(1, 5) 
	res3 := Multiplicar(res1, res2)

	fmt.Println("The value is: ", res3)
} 

func Somar(number1 int, number2 int) int {
	var resultado int = number1 + number2
	return resultado
}

func Subtrair(number1 int ,number2 int ) int {
	var resultado int = number1 - number2
	return resultado
}

func Multiplicar(number1 int ,number2 int ) int {
	var resultado int = number1 * number2
	return resultado
}