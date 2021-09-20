package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var i int
	fmt.Scan(&i)
	list := []int{i}
	var lastElement = list[len(list)-1]
	for lastElement != 2 {
		lastElement = list[len(list)-1]
		if lastElement%2 == 0 {
			var even = lastElement / 2
			list = append(list, even)
		} else {
			var odd = lastElement*3 + 1
			list = append(list, odd)
		}
	}
	fmt.Println(list)
}
