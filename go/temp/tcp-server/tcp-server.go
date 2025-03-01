package main

import (
	"log"
	"net"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in main", r)
		}
	}()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}
		log.Println("Accepted connection from", conn.RemoteAddr())

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading from connection", err)
		return
	}

	_, err = conn.Write(buf)
	if err != nil {
		log.Println("Error writing to connection", err)
		return
	}
}
