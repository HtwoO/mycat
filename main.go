package main

import (
	"net"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1: // no argument provided
		// usage()
	}
	port := ":1111"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		// handle error
		println("Error listening on ", port)
	}
	println("Listening on", port)
	for {
		println("Waiting for connection at", port)
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			println("Error accepting connection ...")
		}
		println("Do nothing with connection at RAM address", conn)
	}
}

func usage() {
	println(os.Args[0], "Usage 使用说明")
}
