package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)
}
