package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var myNumber int
	fmt.Scan(&myNumber)
	list := []int{myNumber}
	var lastElement = myNumber
	for lastElement != 2 {
		lastElement = list[len(list)-1]
		var calc int
		if lastElement%2 == 0 {
			calc = lastElement / 2
		} else {
			calc = lastElement*3 + 1
		}
		list = append(list, calc)
	}
	fmt.Println(list)
}
