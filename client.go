package main

import "net"

func main() {
	srv := "localhost:1111"
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		// handle error
		println("Error connecting to", srv)
	} else {
		println("Connected to", srv)
	}
	err = conn.Close()
	if err != nil {
		println("Error disconnecting to", srv)
	} else {
		println("Connection to", srv, "closed")
	}
}
