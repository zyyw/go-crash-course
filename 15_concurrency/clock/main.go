package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	RunClock()
}

func RunClock() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		go handleConn(conn)
		// handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
