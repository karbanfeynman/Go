package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x+y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	f := 3.1415962
	var z int = int(f)

	fmt.Println("Pi number is", math.Pi)
	fmt.Println(add(4,3))

	fmt.Println(a, b)
	fmt.Println(z)
}
