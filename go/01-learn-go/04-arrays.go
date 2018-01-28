package main

import "fmt"

func main() {
	// An array has a length and a member type
	var a [5]int
	fmt.Println(a)

	// Use square brackets to index an array
	a[2] = 7
	fmt.Println(a)

	// Initialise an array with curly brackets
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	// Remove the element count to create a slice
	s := []int{5, 4, 3, 2, 1}
	fmt.Println(s)

	// Use append to add something new to the array
	s = append(s, 13)
	fmt.Println(s)
}
