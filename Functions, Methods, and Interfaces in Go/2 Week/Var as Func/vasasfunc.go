package main

import "fmt"

//Declara variavel com um padrão de função
var funcVar func(int) int

//Cria a função
func incFn(x int) int {
	return x + 1
}

func main() {
	//atribui a função a variavavel
	funcVar = incFn
	//pode usar a variavel como função
	fmt.Print(funcVar(1))

}
