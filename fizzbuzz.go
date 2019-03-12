package main

import (
	"fmt"
	"strconv"
	"time"
)

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
	fmt.Println("You should now have a 2 second pause before the code resumes")
	time.Sleep(2000 * time.Millisecond)
	userInput()
	tester()
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
