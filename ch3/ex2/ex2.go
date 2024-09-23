package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fourth := runes[3]
	fmt.Println(string(fourth))
}
