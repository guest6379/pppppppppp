package main

import (
	"fmt"
)

func main() {
	fmt.Println("good")

	x := 1
	defer fmt.Println("x", x)
	x = 2

	y := 1
	defer func() {
		fmt.Println("y", y)
	}()
	y = 2

	z := 1
	defer func(z int) {
		fmt.Println("z", z)
	}(z)
	z = 2

	a := 1
	{
		a := a
		defer func() {
			fmt.Println("a", a)
		}()
	}
	a = 2

	b := 1
	defer func() {
		b := b
		fmt.Println("b", b)
	}()
	b = 2
}
