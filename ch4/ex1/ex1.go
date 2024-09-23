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
	fmt.Println(nums)
}
