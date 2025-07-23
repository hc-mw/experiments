package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	os.Args = append(os.Args, "localhost:8080")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET /hello HTTP/1.1\r\nHost: localhost\r\n\r\n")

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

}
