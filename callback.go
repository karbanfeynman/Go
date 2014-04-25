package main

import (
	"fmt"
)

/* type define a function called myFn */
type myFn func(int) string

func foo(i int) string {
	return "hi"
}

/* set function as a parameter */
func bar(fn myFn) string {
	return fn(1)
}

func main() {
	r := bar(foo)
	fmt.Println(r)

	r = bar(func(i int) string {
		return "hello"
	})
	fmt.Println(r)
}
