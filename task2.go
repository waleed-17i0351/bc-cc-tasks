package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	position string
}
type Company struct {
	companyName string
	employees   []Employee
}

// func printP(p Person) {
// 	fmt.Printf("%s is %v years old", p.name, p.age)
// }

func main() {
	fmt.Printf("Hello, World. \n")

	items := []int{1, 3, 4, 61, 11, 0}
	fmt.Println(items)

	e1 := Employee{"Ali", 5000, "Frontend"}
	e2 := Employee{"Senor", 99999, "CEO"}
	e3 := Employee{"Pink", 5050, "Backend"}

	e := []Employee{e1, e2, e3}

	var exsom Company
	exsom.companyName = "ExSoM"
	exsom.employees = e

	fmt.Printf("%v has following employees: %v", exsom.companyName, exsom.employees)
}
