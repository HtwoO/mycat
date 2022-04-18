package main

import (
	"fmt"
	"net"
)

func main() {
	srv := "localhost:1111"
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		// handle error
		fmt.Printf("Error connecting to %s\n", srv)
	} else {
		fmt.Printf("Connected to %s\n", srv)
	}
	err = conn.Close()
	if err != nil {
		fmt.Printf("Error disconnecting to %s\n", srv)
	} else {
		fmt.Printf("Connection to %s closed\n", srv)
	}
}
