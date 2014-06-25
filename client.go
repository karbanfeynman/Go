package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

const (
	MAX_CONCURRENCY = 1024
)

func handle_Connection(conn net.Conn, counter int) {
	_, err_w := conn.Write([]byte("hello"))
	if err_w != nil {
		fmt.Println("Err:", err_w.Error())
		conn.Close()
	}

	buf := make([]byte, 1024)
	_, err_r := conn.Read(buf)
	if err_r != nil {
		fmt.Println("Err:", err_r.Error())
		conn.Close()
	}
	fmt.Printf("[session:%d] %s\n", counter, buf)
	conn.Close()
}

func session(concurrency int) {
	counter := 0
	for {
		conn, err := net.Dial("tcp", "192.168.1.109:8080")

		if err != nil {
			fmt.Printf("[Client:%d]Connection fail\n", counter)
			os.Exit(1)
		} else {
			go handle_Connection(conn, counter)
		}

		time.Sleep(10)
		counter++
		if counter > concurrency {
			break
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()
	session(MAX_CONCURRENCY)
	fmt.Printf("execution time: %s\n", time.Since(start))
	return
}
