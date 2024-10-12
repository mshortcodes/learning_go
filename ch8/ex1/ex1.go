package main

import "fmt"

type num interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~float32 | ~float64
}

func doubler[T num](myNum T) T {
	return myNum * 2
}

func main() {
	fmt.Println(doubler(3))
	fmt.Println(doubler(7.5))
}
