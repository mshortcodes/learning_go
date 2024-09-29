package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{
			FirstName: "Billy",
			LastName:  "Bob",
			Age:       18,
		})
	}
	fmt.Println(people[9_999_999])
}
