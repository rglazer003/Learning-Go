package main

import "fmt"
import "strconv"

func main() {
	for i := 1; i <= 100; i++ {
		var hold string = ""
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
}
