package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* The input is array */
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
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
	}()
	return out
}

/* merge two channels into one */
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	/* set up CPU numbers */
	runtime.GOMAXPROCS(runtime.NumCPU())

	in := gen(2, 3, 10)

	/* two goroutines to share the same channel */
	c1 := sq(in)
	c2 := sq(in)
	for out := range merge(c1, c2) {
		fmt.Println(out)
	}

}
