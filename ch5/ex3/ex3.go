package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return prefix + " " + body
	}
}

func main() {
	prefix := prefixer("hello from prefix")
	fmt.Println(prefix("hello from body"))
}
