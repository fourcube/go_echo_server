// A simple echo server written in go
package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:4040")
	if err != nil {
		log.Fatalf("Couldn't resolve TCP address: %v", err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Couldn't listen on %v TCP address: %v", addr, err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatalf("Couldn't accept on %v TCP address: %v", addr, err)
		}

		go handle(conn)
	}
}

func handle(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("Couldn't read from tcp socket: %v", err)
			return
		}

		bytesWritten, err := conn.Write(data)
		if err != nil {
			log.Printf("Failed write on tcp socket: %v", err)
			return
		}

		if bytesWritten != len(data) {
			log.Printf("Warning: Incomplete write on tcp socket")
			return
		}
	}

}
