package main

import "fmt"

type A struct {
	x int
	y int
}

func main() {
	a := A{x: 1}
	b := &a

	b.x = 4

	fmt.Println(a)
}
