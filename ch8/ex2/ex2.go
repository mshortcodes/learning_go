package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintNum[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var i PrintInt = 10
	PrintNum(i)

	var f PrintFloat = 10.5
	PrintNum(f)
}
