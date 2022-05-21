package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1: // no argument provided
		usage()
	}
	port := ":1111"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		// handle error
		fmt.Printf("Error listening on %s\n", port)
	}
	fmt.Printf("Listening on %s\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Error accepting connection ...")
		}
		defer func() {
			// Shut down the connection
			fmt.Printf("Closing connection: %T\n", conn)
			conn.Close()
			fmt.Printf("Connection closed: %q\n", conn)
		}()
		fmt.Printf("Connection received from: %s\n", conn.RemoteAddr())
		// Receive data (write incoming data to os.Stdout)
		go func(conn net.Conn) {
			for {
				if _, err = io.Copy(os.Stdout, conn); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}(conn)
		// Send data (read data from os.Stdin)
		go func(conn net.Conn) {
			for {
				if _, err = io.Copy(conn, os.Stdin); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}(conn)
	}
}

func usage() {
	fmt.Printf("%s, Usage 使用说明\n", os.Args[0])
}
