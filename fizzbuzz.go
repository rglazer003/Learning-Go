package main

import "fmt"
import "strconv"

func main() {
	for i := 1; i <= 100; i++ {
		var hold string
		if i%3 == 0 {
			hold += "Fizz"
		}
		if i%5 == 0 {
			hold += "Buzz"
		}
		if i%3 != 0 && i%5 != 0 {
			hold = strconv.Itoa(i)
		}
		fmt.Println(hold)
	}
	exp()
	userInput()
}
func exp() {
	fmt.Println("And that's fizzbuzz!")
}
func userInput() {
	fmt.Println("What number would you like to see the square of?")
	var number int
	fmt.Scan(&number)
	var result = number * number
	fmt.Printf("%v squared is %v", number, result)
	fmt.Println()
}
