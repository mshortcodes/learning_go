package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		num := rand.Intn(100)
		nums = append(nums, num)
	}
loop:
	for _, num := range nums {
		switch {
		case num%2 == 0 && num%3 == 0:
			fmt.Println("Six!")
			break loop
		case num%2 == 0:
			fmt.Println("Two!")
		case num%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
