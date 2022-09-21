package main

import "fmt"

type Person struct {
	name string
	age  int
}

func printP(p Person) {
	fmt.Printf("%s is %v years old", p.name, p.age)
}

func main() {
	fmt.Printf("Hello, World. \n")

	items := []int{1, 3, 4, 61, 11, 0}
	fmt.Println(items)

	var p1 Person
	p1.name = "James"
	p1.age = 25

	printP(p1)
}
