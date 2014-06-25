package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)

	fmt.Printf("A connection establishes!\n")
	fmt.Println(conn.RemoteAddr())

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	conn.Write([]byte("msg received!"))
	conn.Close()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("fail to create tcp listening\n")
		os.Exit(1)
		return
	}

	/* Close the listener when the app closes */
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Connection fail\n")
			os.Exit(1)
			continue
		} else {
			go handleConnection(conn)
		}
	}
}
