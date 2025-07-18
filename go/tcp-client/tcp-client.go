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

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}

	log.Println("Received:", string(buf))
}
