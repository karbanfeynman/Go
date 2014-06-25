package main

import (
	"fmt"
	"runtime"
	"time"
)

/* The input is array */
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		fmt.Printf("sleep..\n")
		time.Sleep(5)
		close(out)
		fmt.Printf("close channel\n")
	}()
	return out
}

/* input is a channel, and out is another channel */
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
		fmt.Printf("[sq]channel close\n")
	}()
	return out
}

func main() {
	/* set up CPU numbers */
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := gen(2, 3, 10)
	for out := range sq(c) {
		fmt.Println(out)
	}

}
