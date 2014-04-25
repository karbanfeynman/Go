package main

import (
	"fmt"
	"runtime"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/* Main function */
func main() {
	numCPU := runtime.NumCPU()
	c := make(chan int, 10)

	fmt.Println("number of CPU: ", numCPU)
	if numCPU > 1 {
		runtime.GOMAXPROCS(numCPU)
	}
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
