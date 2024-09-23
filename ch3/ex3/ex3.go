package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := Employee{
		"Billy",
		"Bill",
		0,
	}

	e2 := Employee{
		firstName: "Bobby",
		lastName:  "Bob",
		id:        1,
	}

	var e3 Employee
	e3.firstName = "Johnny"
	e3.lastName = "John"
	e3.id = 2

	fmt.Println(e1, e2, e3)
}
