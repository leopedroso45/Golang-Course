package main

import "fmt"

func BubbleSort(slice []int) {
	fmt.Println(len(slice))
	var finish int = 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > slice[i+1] {
			finish++
			Swap(slice, i)
			println(slice)
		} else {
			finish--
		}
	}
}

func Swap(slice []int, i int) error {

	temp := slice[i+1]
	slice[i+1] = slice[i]
	slice[1] = temp
}

/*GetNumbers get ten numbers from the user and it returns a slice*/
func GetNumbers() []int {
	var intReceived int
	var numbersReceived []int

	for i := 0; i < 10; i++ {
		fmt.Printf("Type an integer please:")
		valor, _ := fmt.Scan(&intReceived)
		if valor != 0 {
			numbersReceived = append(numbersReceived, intReceived)
		} else {
			i--
		}
	}
	return numbersReceived
}

func main() {

	firstSlice := GetNumbers()
	fmt.Println(firstSlice)
	BubbleSort(firstSlice)

}
