package main

import "fmt"

func main() {
	x := 5

	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("Something else")
	}
}
