/*
If you use infinite loop, CPU usage will be 100% all the time.
*/
package main

import (
	"runtime"
	"time"
)

func loop() {
	for {
		time.Sleep(10)
	}
}
func main() {
	number := runtime.NumCPU()
	runtime.GOMAXPROCS(number)
	go loop()
	loop()
}
