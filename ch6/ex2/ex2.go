package main

import "fmt"

func UpdateSlice(strs []string, str string) {
	strs[len(strs)-1] = str
	fmt.Printf("In UpdateSlice: %s\n", strs)
}

func GrowSlice(strs []string, str string) {
	strs = append(strs, str)
	fmt.Printf("In GrowSlice: %s\n", strs)
}

func main() {
	strs := []string{"hello", "how", "are", "you"}
	UpdateSlice(strs, "Bob")
	fmt.Printf("In main after UpdateSlice: %s\n", strs)
	GrowSlice(strs, "John")
	fmt.Printf("In main after GrowSlice: %s\n", strs)
}
