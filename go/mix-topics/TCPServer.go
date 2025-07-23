package main

import (
	"log"
	"net"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func serverConRequest(conn net.Conn) {
	// close connection at the end
	defer conn.Close()
	// read (syscall-blocking) the connection
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	handleError(err)
	time.Sleep(1 * time.Second)
	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World\r\n"))
	handleError(err)
}

func TCPListen() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	// step 1: listen to a port for tcp connection
	listener, err := net.Listen("tcp", ":3000")
	// handle error
	handleError(err)

	log.Println("TCP Server Listening on PORT:3000...")

	for {
		// accept (syscall) connection (blocking: waits till a connection is made)
		conn, err := listener.Accept()
		// handle error
		handleError(err)
		log.Println("Connection From: ", conn.RemoteAddr())
		// serve the connection request
		go serverConRequest(conn)
	}
}
