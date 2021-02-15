package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "\nHello from Andrew's Laptop\n")
		fmt.Fprintln(conn, "How was your day?")
		fmt.Fprintf(conn, "%v", "Well I hope!\n")
		conn.Close()
	}
}
